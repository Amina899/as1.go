package main

import (
	"aitu/aitunews/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	n, err := app.news.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Newses: n,
	})
}

func (app *application) showNews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	n, err := app.news.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "show.page.tmpl", &templateData{
		News: n,
	})
}

func (app *application) createNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "Think Tank Thrive - Ideathon in Industry Quandaries'23"
	author := "Zhansaya Makhambetova"
	content := "29-30 ноября пройдет международный хакатон\n- Think Tank Thrive - Ideathon in Industry Quandaries'23.\n\nОрганизаторами являются Astana IT University, Woxsen University,\nAlmaty Management University.\n\nЦелью Ideathon является генерирование творческих идей, выявление возможностей \nи поощрение студентов к налаживанию связей и общению с наставниками и лидерами отрасли.\n\nФиналистам будут начислены по 25-30 баллов ROS GPA каждому студенту,\nтакже всем участникам будут вручены сертификаты от Woxsen University (Индия)\n\nФормат мероприятия: Онлайн\n\nРегистрация открыта до 20 ноября (включительно). https://clck.ru/36WNhj\n\n✔️ Подробная информация доступна на https://ideathon.aircwou.in/"
	created := "Monday, 13 November 2023, 10:27 AM"
	_, err := time.Parse("24 Jan 2024 at 04:23", created)
	id, err := app.news.Insert(title, content, author, created)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/news?id=%d", id), http.StatusSeeOther)
}
func (app *application) category(writer http.ResponseWriter, request *http.Request) {

}
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	app.render(w, r, "contact.page.tmpl", nil)
}
