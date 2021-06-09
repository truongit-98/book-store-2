package main

import (
	"BookStore/common/structs"
	"BookStore/models"
	"BookStore/utils"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)
func readFile(url string) {

}

func init()  {
	_, _ = models.InitDB()

}

func main(){
	//fakeDataCsv()
	//fakeDataUser()
	fakeOrder()
	fakeOrderDetail()
	//updatePublisher()
}

func fakeDataUser()  {

	address := []string{"Hà Nội", "Thanh Hoá", "Bắc Ninh", "Vĩnh Phúc", "TP Hồ Chí Minh", "Hải Phnòg", "Thái Bình", "Nghệ An"}
	for i:=1; i<=100;i++{
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 7
		addressRand := rand.Intn(max-min+1) + min
		user := models.Customer{
			Email: fmt.Sprintf("truong%d@gmail.com", i),
			Password: "truong123",
			FullName: fmt.Sprintf("Nguyễn Văn A"),
			Address: address[addressRand],
		}

		models.Create(&user)
	}

}

func fakeOrder()  {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-10-11T11:45:26.371Z"
	t, _ := time.Parse(layout, str)
	fmt.Println(t.Month())
	rand.Seed(time.Now().UnixNano())

	items, _ := models.GetAll(&models.Customer{})
	users := items.([]models.Customer)
	users = users[0:30]
	for _, u := range users {
		for i :=0;i <3; i++{
			min := 1
			max := 10
			day := rand.Intn(max-min+1) + min
			crated := t

			obj := models.Order{
				DateCreated: crated.Add((time.Hour * 24) * time.Duration(day)).Unix(),
				Status: "Đang chờ xử lý",
				OrderDate: crated.Unix(),
				CustomerID: &u.ID,
			}
			models.Create(&obj)
		}
	}
}

func fakeOrderDetail()  {
	items, _ := models.GetAll(&models.Order{})
	orders := items.([]models.Order)

	for _, o := range orders{
		for i:=0;i<3;i++{
			items, _ := models.GetAll(models.Book{})
			books := items.([]models.Book)
			rand.Seed(time.Now().UnixNano())
			min := 10
			max := 120
			amount := rand.Intn(max-min+1) + min
			obj := models.OrderDetail{
				OrderID: o.ID,
				BookID: books[amount].ID,
				Quantity: int(amount/10),
				Price: float64(amount) * 3 * 30000,
			}
			models.Create(&obj)
		}
	}
}

func updatePublisher()  {

	items ,_ := models.GetAll(models.BookDetail{})
	details := items.([]models.BookDetail)
	for _, d := range details{
		rand.Seed(time.Now().UnixNano())
		min := 2
		max := 20
		topicId := rand.Intn(max-min+1) + min
		d.PublisherID = uint(topicId)
		models.Update(d)
	}
}

func fakeDataCsv()  {
	csvFile, err := os.Open( strings.Join([]string{utils.GetCurrentPath(), "/tools/book.csv"},""))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true
	csvLines, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := make([]structs.BookDataFake, 0)
	csvLines = csvLines[1:]
	for i, line := range csvLines {
		datas := strings.Split(string(line[0]), ",")
		if len(datas) == 12 && i<=1000 {
			obj := structs.BookDataFake{
				Title: datas[1],
				Authors:  datas[2],
				Average_rating:  datas[3],
				Isbn:  datas[4],
				Isbn13:  datas[5],
				Language_code:  datas[6],
				Num_pages:  datas[7],
				Ratings_count:  datas[8],
				Text_reviews_count:  datas[9],
				Publication_date:  datas[10],
				Publisher:  datas[11],
			}
			result = append(result, obj)
		}
	}
	for idx, job := range result{
		bType := models.BookType{
			Type:  "Non series",
			TypeName: job.Title,
			Episodes: 1,
		}
		bTypeId, _ := bType.Create2()
		layout := "2006-01-02T15:04:05.000Z"
		str := "2020-10-11T11:45:26.371Z"
		t, _ := time.Parse(layout, str)
		rand.Seed(time.Now().UnixNano())
		book := models.Book{
			Title: job.Title,
			ISBN: job.Isbn,
			Ratting: parseFloat(job.Average_rating),
			BookTypeID: &bTypeId,
			CreatedAt: t.Add(time.Hour * 24 * time.Duration(idx) *3).Unix(),
		}

		book.Create2()
		authorId, _ := (models.Author{
			AuthorName: job.Authors,
		}).Create2()
		(models.BookTypeAuthor{
			BookTypeID: bTypeId,
			AuthorID: authorId,
		}).Create()
		publisherId, _ := (models.Publisher{
			PublisherName: job.Publisher,
		}).Create2()
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 7
		topicId := rand.Intn(max-min+1) + min
		price := float64(rand.Intn(max-min+1) + min) * 30000
		language := []string{"en-US", "vn", "eng", "en-US", "vn", "vn", "eng","eng"  }
		detail := models.BookDetail{
			Price: price,
			PriceCover: price,
			PublisherYear: int(parseInt(strings.SplitAfter(job.Publication_date, "/")[2])),
			NumberOfPage: int(parseInt(job.Num_pages)),
			Height: 20,
			Width: 14,
			Description: "Đây là năm thứ 6 của Harry Potter tại trường Hogwarts. Trong lúc những thế lực hắc ám của Voldemort gieo rắc nỗi kinh hoàng và sợ hãi ở khắp nơi, mọi chuyện càng lúc càng trở nên rõ ràng hơn đối với Harry, rằng cậu sẽ sớm phải đối diện với định mệnh của mình. Nhưng liệu Harry đã sẵn sàng vượt qua những thử thách đang chờ đợi phía trước?",
			Language: language[topicId],
			PublisherID: publisherId,
			BookTypeID: bTypeId,
			FormatID: 1,
			TopicID: uint(topicId),
		}
		models.Create(detail)
	}
	log.Println("Okkkkkkkkkkkkkkkkkk")
}

func parseFloat( s string) float64 {
	r, _ := strconv.ParseFloat(s, 64)
	return r
}

func parseInt(s string) int32 {
	r, _ := strconv.ParseInt(s, 10, 32)
	return int32(r)
}