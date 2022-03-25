package book

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func GetBooksWithWorkerPool(path string) ([]Book, error) {
	const numJobs = 5
	books := []Book{}
	jobs := make(chan []string, numJobs)
	results := make(chan Book, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 3; w++ {
		fmt.Println("Worker starting", w)
		wg.Add(1)
		go toStruct(jobs, results, &wg)
	}
	go func() {
		fmt.Println("open file running...")
		f, _ := os.Open(path)
		defer f.Close()

		lines, _ := csv.NewReader(f).ReadAll()

		for _, line := range lines[1:] {
			fmt.Println("line", line[0])
			jobs <- line
		}
		close(jobs)
	}()

	go func() {
		fmt.Println("wait")
		wg.Wait()
		close(results)
	}()

	fmt.Println(results)
	for b := range results {
		books = append(books, b)
	}
	return books, nil
}

func toStruct(jobs <-chan []string, results chan<- Book, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		pageNumberParsed, _ := strconv.Atoi(j[2])
		// if err != nil {
		// 	return err
		// }
		stockNumberParsed, _ := strconv.Atoi(j[3])
		// if err != nil {
		// 	return err
		// }
		priceParsed, _ := strconv.ParseFloat(j[5], 0)
		// if err != nil {
		// 	return err
		// }

		isDeletedParsed, _ := strconv.ParseBool(j[7])
		// if err != nil {
		// 	return err
		// }
		book := Book{ID: j[0],
			Name:        j[1],
			PageNumber:  uint(pageNumberParsed),
			StockNumber: stockNumberParsed,
			StockID:     j[4],
			Price:       float32(priceParsed),
			ISBN:        j[6],
			IsDeleted:   isDeletedParsed,
			AuthorID:    j[8],
			AuthorName:  j[9]}

		results <- book

	}

}
