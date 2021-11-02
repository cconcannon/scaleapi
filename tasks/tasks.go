package tasks

type Task struct {
	TaskId            string                 `json:"task_id"`
	CreatedAt         string                 `json:"created_at"`
	UpdatedAt         string                 `json:"updated_at"`
	CompletedAt       string                 `json:"completed_at"`
	Type              string                 `json:"type"`
	Status            string                 `json:"status"`
	Instruction       string                 `json:"instruction"`
	Params            Params                 `json:"params"`
	CallbackUrl       string                 `json:"callback_url"`
	CallbackCompleted bool                   `json:"callback_completed"`
	Response          Response               `json:"response"`
	Metadata          map[string]interface{} `json:"metadata"`
	Audits            []Audit                `json:"audits"`
	Tags              []string               `json:"tags"`
	UniqueId          string                 `json:"unique_id"`
}

type Params struct {
	AttachmentType string   `json:"attachment_type"`
	Attachment     string   `json:"attachment"`
	Categories     []string `json:"categories"`
}

type Response struct {
	Category string `json:"category"`
}

type Audit struct {
	AuditedBy     string `json:"audited_by"`
	AuditedAt     string `json:"audited_at"`
	AuditTimeSecs int    `json:"audit_time_secs"`
	AuditResult   string `json:"audit_result"`
	AuditSource   string `json:"audit_source"`
}
