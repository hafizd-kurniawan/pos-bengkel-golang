package models

import (
	"time"
)

// ServiceJobs table
type ServiceJob struct {
	ServiceJobID               uint              `db:"service_job_id" json:"service_job_id"`
	ServiceCode                string            `db:"service_code" json:"service_code"`
	QueueNumber                int               `db:"queue_number" json:"queue_number"`
	CustomerID                 uint              `db:"customer_id" json:"customer_id"`
	VehicleID                  uint              `db:"vehicle_id" json:"vehicle_id"`
	TechnicianID               *uint             `db:"technician_id" json:"technician_id"`
	ReceivedByUserID           uint              `db:"received_by_user_id" json:"received_by_user_id"`
	OutletID                   uint              `db:"outlet_id" json:"outlet_id"`
	ProblemDescription         string            `db:"problem_description" json:"problem_description"`
	TechnicianNotes            *string           `db:"technician_notes" json:"technician_notes"`
	Status                     ServiceStatusEnum `db:"status" json:"status"`
	ServiceInDate              time.Time         `db:"service_in_date" json:"service_in_date"`
	PickedUpDate               *time.Time        `db:"picked_up_date" json:"picked_up_date"`
	ComplainDate               *time.Time        `db:"complain_date" json:"complain_date"`
	WarrantyExpiresAt          *time.Time        `db:"warranty_expires_at" json:"warranty_expires_at"`
	NextServiceReminderDate    *time.Time        `db:"next_service_reminder_date" json:"next_service_reminder_date"`
	DownPayment                float64           `db:"down_payment" json:"down_payment"`
	GrandTotal                 float64           `db:"grand_total" json:"grand_total"`
	TechnicianCommission       float64           `db:"technician_commission" json:"technician_commission"`
	ShopProfit                 float64           `db:"shop_profit" json:"shop_profit"`
	CreatedAt                  time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt                  time.Time         `db:"updated_at" json:"updated_at"`
	DeletedAt                  *time.Time        `db:"deleted_at" json:"deleted_at"`
	CreatedBy                  *uint             `db:"created_by" json:"created_by"`

	// Relationships (populated separately)
	Customer       *Customer            `json:"customer,omitempty"`
	Vehicle        *CustomerVehicle     `json:"vehicle,omitempty"`
	Technician     *User                `json:"technician,omitempty"`
	ReceivedByUser *User                `json:"received_by_user,omitempty"`
	Outlet         *Outlet              `json:"outlet,omitempty"`
	ServiceDetails []ServiceDetail      `json:"service_details,omitempty"`
	Histories      []ServiceJobHistory  `json:"histories,omitempty"`
}

// ServiceDetails table
type ServiceDetail struct {
	DetailID        uint    `db:"detail_id" json:"detail_id"`
	ServiceJobID    uint    `db:"service_job_id" json:"service_job_id"`
	ItemID          uint    `db:"item_id" json:"item_id"`
	ItemType        string  `db:"item_type" json:"item_type"`
	Description     string  `db:"description" json:"description"`
	SerialNumberUsed *string `db:"serial_number_used" json:"serial_number_used"`
	Quantity        int     `db:"quantity" json:"quantity"`
	PricePerItem    float64 `db:"price_per_item" json:"price_per_item"`
	CostPerItem     float64 `db:"cost_per_item" json:"cost_per_item"`

	// Relationships (populated separately)
	ServiceJob *ServiceJob `json:"service_job,omitempty"`
}

// ServiceJobHistories table
type ServiceJobHistory struct {
	HistoryID    uint       `db:"history_id" json:"history_id"`
	ServiceJobID uint       `db:"service_job_id" json:"service_job_id"`
	UserID       uint       `db:"user_id" json:"user_id"`
	Notes        *string    `db:"notes" json:"notes"`
	ChangedAt    time.Time  `db:"changed_at" json:"changed_at"`

	// Relationships (populated separately)
	ServiceJob *ServiceJob `json:"service_job,omitempty"`
	User       *User       `json:"user,omitempty"`
}