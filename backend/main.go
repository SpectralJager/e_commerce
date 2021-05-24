package main

type User struct {
	Id       string
	Username string
	Password string
	BoxId    string
}

type Token struct {
	Id     string
	UserId string
	JWT    string
}

type Product struct {
	Id   string
	Name string
}

type Box struct {
	Id         string
	UserId     string
	ProductsId string
}

var (
	Users    []User
	Products []Product
	Tokens   []Token
	Boxes    []Box
)

func main() {
	HandleRequests()
}

/*
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var articles []Article

func main() {
	articles = []Article{
		Article{Id: "1", Title: "First article", Desc: "its my first article, check it!", Content: "Minim do incididunt est aute occaecat id aute velit aliqua eu dolor laborum magna proident. Amet ullamco est culpa in consectetur. Culpa minim nostrud ipsum qui. Aliquip fugiat esse exercitation duis ea id ullamco ea nisi voluptate minim laboris et. Irure deserunt aliqua aliquip eu non pariatur est et ea mollit."},
		Article{Id: "2", Title: "Second and last article", Desc: "I'm disapointment", Content: "Lorem labore sint ex anim aliqua incididunt nostrud."},
	}
	handleRequests()
}

func handleRequests() {
	multiplexer := mux.NewRouter().StrictSlash(true)
	multiplexer.HandleFunc("/articles", ReturnAllArticles)
	multiplexer.HandleFunc("/articles/{id}", ReturnSingleArticle).Methods("GET")
	multiplexer.HandleFunc("/articles/{id}", CreateArticle).Methods("POST")
	multiplexer.HandleFunc("/articles/{id}", DeleteArticle).Methods("DELETE")
	multiplexer.HandleFunc("/articles/{id}", UpdateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", multiplexer))
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: ReturnSingleArticle")
	w.Header().Add("Content-type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			break
		}
	}

}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: ReturnAllArticles")
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: CreateArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: DeleteArticle")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: UpdateArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var temp_article Article
	json.Unmarshal(reqBody, &temp_article)
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range articles {
		if article.Id == id {
			temp := articles[index+1:]
			articles = append(articles[:index], temp_article)
			articles = append(articles[:index+1], temp...)
			break
		}
	}
} */
