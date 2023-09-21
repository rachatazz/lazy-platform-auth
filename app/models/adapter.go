package model

type SendMailWithTemplateRequest struct {
	Template    string
	SendTo      string
	DisplayName string
	Ticket      string
}

type TemplateBodyResponse struct {
	Success *bool   `json:"success,omitempty"`
	Data    *Data   `json:"data,omitempty"`
	Error   *string `json:"error,omitempty"`
}

type Data struct {
	Templates            []TemplateResponse `json:"templates,omitempty"`
	Templatescount       *int64             `json:"templatescount,omitempty"`
	Drafttemplate        []interface{}      `json:"drafttemplate,omitempty"`
	Drafttemplatescount  *int64             `json:"drafttemplatescount,omitempty"`
	Globaltemplate       []TemplateResponse `json:"globaltemplate,omitempty"`
	Globaltemplatescount *int64             `json:"globaltemplatescount,omitempty"`
	Parenttemplate       []interface{}      `json:"parenttemplate,omitempty"`
	Parenttemplatescount *int64             `json:"parenttemplatescount,omitempty"`
}

type TemplateResponse struct {
	Templateid           *int64      `json:"templateid,omitempty"`
	Templatetype         *int64      `json:"templatetype,omitempty"`
	Name                 *string     `json:"name,omitempty"`
	Editormetadata       *string     `json:"editormetadata,omitempty"`
	Dateadded            *string     `json:"dateadded,omitempty"`
	CSS                  interface{} `json:"css,omitempty"`
	Subject              *string     `json:"subject,omitempty"`
	Fromemail            *string     `json:"fromemail,omitempty"`
	Fromname             *string     `json:"fromname,omitempty"`
	Bodyhtml             interface{} `json:"bodyhtml,omitempty"`
	Bodyamp              interface{} `json:"bodyamp,omitempty"`
	Bodytext             interface{} `json:"bodytext,omitempty"`
	Originaltemplateid   *int64      `json:"originaltemplateid,omitempty"`
	Originaltemplatename *string     `json:"originaltemplatename,omitempty"`
	Templatescope        *int64      `json:"templatescope,omitempty"`
	Screenshot           *string     `json:"screenshot,omitempty"`
	Thumbnail            *string     `json:"thumbnail,omitempty"`
	Mediumscreenshot     *string     `json:"mediumscreenshot,omitempty"`
	Ispublic             *bool       `json:"ispublic,omitempty"`
	Isdefault            *bool       `json:"isdefault,omitempty"`
	Link                 *string     `json:"link,omitempty"`
	Tags                 []string    `json:"tags,omitempty"`
}
