package scaleapi

type TasksResponse struct {
	Tasks   []Task `json:"docs"`
	Total   int    `json:"total"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	HasMore bool   `json:"has_more"`
}

type Task struct {
	TaskId                 string                 `json:"task_id"`
	CreatedAt              string                 `json:"created_at"`
	UpdatedAt              string                 `json:"updated_at"`
	CompletedAt            string                 `json:"completed_at"`
	Type                   string                 `json:"type"`
	Status                 string                 `json:"status"`
	Instruction            string                 `json:"instruction"`
	Params                 Params                 `json:"params"`
	IsTest                 bool                   `json:"is_test"`
	Urgency                string                 `json:"urgency"`
	ProcessedAttachments   []interface{}          `json:"processed_attachments"`
	Project                string                 `json:"project"`
	CustomerReviewStatus   string                 `json:"customer_review_status"`
	CustomerReviewComments []string               `json:"customer_review_comments"`
	PriorResponses         []interface{}          `json:"prior_responses"`
	CallbackUrl            string                 `json:"callback_url"`
	CallbackCompleted      bool                   `json:"callback_completed"`
	Response               Response               `json:"response"`
	CustomerReviewedBy     string                 `json:"customer_reviewed_by"`
	CustomerAuditTimeSecs  float64                `json:"customer_audit_time_secs"`
	CustomerAuditedAt      string                 `json:"customer_audited_at"`
	Metadata               map[string]interface{} `json:"metadata"`
	Audits                 []Audit                `json:"audits"`
	Tags                   []string               `json:"tags"`
	UniqueId               string                 `json:"unique_id"`
}

type Params struct {
	AttachmentType       string               `json:"attachment_type"`
	Attachment           string               `json:"attachment"`
	ObjectsToAnnotate    []string             `json:"objects_to_annotate"`
	WithLabels           bool                 `json:"with_labels"`
	MinWidth             int                  `json:"min_width"`
	MinHeight            int                  `json:"min_height"`
	Examples             []interface{}        `json:"examples"`
	AnnotationAttributes AnnotationAttributes `json:"annotation_attributes"`
}

type AnnotationAttributes struct {
	Occlusion       Occlusion       `json:"occlusion"`
	Truncation      Truncation      `json:"truncation"`
	BackgroundColor BackgroundColor `json:"background_color"`
}

type Occlusion AnnotationAttribute

type Truncation AnnotationAttribute

type BackgroundColor AnnotationAttribute

type AnnotationAttribute struct {
	Description string   `json:"description"`
	Choices     []string `json:"choices"`
}

type Response struct {
	Annotations      []Annotation           `json:"annotations"`
	GlobalAttributes map[string]interface{} `json:"global_attributes"`
	IsCustomerFix    bool                   `json:"is_customer_fix"`
}

type Annotation struct {
	Label      string       `json:"label"`
	Attributes AttributeSet `json:"attributes"`
	Uuid       string       `json:"uuid"`
	Width      int          `json:"width"`
	Height     int          `json:"height"`
	Geometry   string       `json:"geometry"`
	Left       float64      `json:"left"`
	Top        float64      `json:"top"`
}

type AttributeSet struct {
	Occlusion       string `json:"occlusion"`
	Truncation      string `json:"truncation"`
	BackgroundColor string `json:"background_color"`
}

type Audit struct {
	AuditedBy           string  `json:"audited_by"`
	AuditedAt           string  `json:"audited_at"`
	AuditTimeSecs       float64 `json:"audit_time_secs"`
	AuditActiveTimeSecs float64 `json:"audit_active_time_secs,omitempty"`
	AuditResult         string  `json:"audit_result"`
	AuditSource         string  `json:"audit_source"`
}
