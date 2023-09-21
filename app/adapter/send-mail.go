package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	model "lazy-platform-auth/app/models"
	"lazy-platform-auth/config"
)

type ElasticEmailAdapter interface {
	SendMailWithTemplate(req model.SendMailWithTemplateRequest) error
	GetEmailTemplate(templateName string) (*model.TemplateResponse, error)
}

type elasticEmailAdapter struct {
	ElasticUrl    string
	ElasticApiKey string
}

func NewElasticEmailAdapter(config config.Config) ElasticEmailAdapter {
	return elasticEmailAdapter{
		ElasticUrl:    config.ElasticUrl,
		ElasticApiKey: config.ElasticApiKey,
	}
}

func (a elasticEmailAdapter) SendMailWithTemplate(req model.SendMailWithTemplateRequest) error {
	formData := url.Values{
		"apikey":             {a.ElasticApiKey},
		"template":           {req.Template},
		"msgTo":              {req.SendTo},
		"merge_display_name": {req.DisplayName},
		"merge_ticket":       {req.Ticket},
	}
	resp, err := http.PostForm(a.ElasticUrl+"/email/send", formData)
	if err != nil {
		return err
	}

	fmt.Printf("%v", resp.Body)
	defer resp.Body.Close()

	return nil
}

func (a elasticEmailAdapter) GetEmailTemplate(
	templateName string,
) (*model.TemplateResponse, error) {
	url := a.ElasticUrl + "/template/getlist"
	params := Params{"apikey": a.ElasticApiKey}
	resp, err := HttpGetRequest(
		url,
		&AdapterOptions{params: params},
	)
	if err != nil {
		return nil, err
	}

	templateBodyResponse := model.TemplateBodyResponse{}
	json.NewDecoder(resp.Body).Decode(&templateBodyResponse)
	defer resp.Body.Close()
	templates := templateBodyResponse.Data.Templates

	if templates != nil {
		template := model.TemplateResponse{}
		for _, value := range templates {
			if *value.Name == templateName {
				template = value
				break
			}
		}
		return &template, nil
	}

	return nil, errors.New("templates not fount")
}
