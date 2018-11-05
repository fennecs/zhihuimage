package service

import (
	"regexp"
	"net/http"
	"zhihuimage/entity"
	"os"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
)

const Url = "https://www.zhihu.com/api/v4/questions/{QUESTION_ID}/answers?include=content&limit={LIMIT}&offset={OFFSET}&sort_by=default"
const ZhihuUrl = "https://www.zhihu.com/question/{QUESTION_ID}"
const Size = 20

func GetWonderfulImages(questionId int, rootDir string) {
	totalCount := getAnswerCount(questionId)
	fmt.Println("the question has total:", totalCount, "answers")
	for ; ;  {

	}
}

func getAnswerCount(questionId int) int {
	response, err := http.Get(strings.Replace(ZhihuUrl, "{QUESTION_ID}", strconv.Itoa(questionId), -1))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	answerCountContent, _ := doc.Find(".App-main").Find("meta[itemProp=answerCount]").Eq(0).Attr("content")
	answerCount, err := strconv.Atoi(answerCountContent)
	if err != nil {
		panic(err)
	}
	return answerCount
}

func downloadImageByApi(url string, rootDir string) {
	results := make([]*entity.Image, 0)
	answers := getPagingAnswer(url)

	for _, answer := range answers.Data {
		//fmt.Println(answer.Content)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(answer.Content))
		if err != nil {
			panic(err)
		}

		doc.Find("figure img").Each(func(i int, selection *goquery.Selection) {
			originImgUrl, exist := selection.Attr("data-original")
			if !exist {
				fmt.Println("不存在原图,可能是个表情包")
				return
			}

			result := make(chan bool)
			image := &entity.Image{
				Path:   originImgUrl,
				Result: result,
			}

			go doDownload(image, rootDir)
			results = append(results, image)
		})
	}

	fmt.Println("All task started, wating for finishment...")

	waitForFinish(results)
}

func getPagingAnswer(url string) entity.PagingAnswer {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))
	var answers entity.PagingAnswer
	err = json.Unmarshal(body, &answers)
	if err != nil {
		panic(err)
	}
	return answers
}

func waitForFinish(results []*entity.Image) {
	for _, image := range results {
		fmt.Println(image.Path, <-image.Result)
		close(image.Result)
	}
}

func doDownload(image *entity.Image, imgDir string) {
	res, err := http.Get(image.Path)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	name := getFileName(image.Path)
	out, err := os.Create(imgDir + name)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(out, res.Body)
	if err != nil {
		image.Result <- false
		panic(err)
	}

	//fmt.Println("[", name ,"]downloaded")

	image.Result <- true
}

func getFileName(path string) string {
	reg, _ := regexp.Compile(`(\w|\d|_)*.jpg`)
	name := reg.FindStringSubmatch(path)[0]
	return name
}
