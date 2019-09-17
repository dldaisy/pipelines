// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"
)

// PipelineVersionStatus a label for the status of the Pipeline.
// This is intend to make pipeline creation and deletion atomic.
type PipelineVersionStatus string

const (
	PipelineVersionCreating PipelineVersionStatus = "CREATING"
	PipelineVersionReady    PipelineVersionStatus = "READY"
	PipelineVersionDeleting PipelineVersionStatus = "DELETING"
)

type PipelineVersion struct {
	UUID           string `gorm:"column:UUID; not null; primary_key"`
	CreatedAtInSec int64  `gorm:"column:CreatedAtInSec; not null; index"`
	Name           string `gorm:"column:Name; not null; unique"`
	// Set size to 65535 so it will be stored as longtext.
	// https://dev.mysql.com/doc/refman/8.0/en/column-count-limit.html
	Parameters string `gorm:"column:Parameters; not null; size:65535"`
	// PipelineVersion belongs to Pipeline. If a pipeline with a specific UUID
	// is deleted from Pipeline table, all this pipeline's versions will be
	// deleted from PipelineVersion table.
	PipelineId string                `gorm:"column:PipelineId; not null;index;"`
	Status     PipelineVersionStatus `gorm:"column:Status; not null"`
	// Code source refers to the pipeline version's definition.
	// We allow multiple code source references for a single pipeline version.
	// CodeSourceURLs stores URL links to code source, separated by ";"
	CodeSourceUrls string `gorm:"column:CodeSourceUrls;"`
}

func (p PipelineVersion) GetValueOfPrimaryKey() string {
	return fmt.Sprint(p.UUID)
}

func GetPipelineVersionTablePrimaryKeyColumn() string {
	return "UUID"
}

// PrimaryKeyColumnName returns the primary key for model PipelineVersion.
func (p *PipelineVersion) PrimaryKeyColumnName() string {
	return "UUID"
}

// DefaultSortField returns the default sorting field for model Pipeline.
func (p *PipelineVersion) DefaultSortField() string {
	return "CreatedAtInSec"
}

// APIToModelFieldMap returns a map from API names to field names for model
// PipelineVersion.
func (p *PipelineVersion) APIToModelFieldMap() map[string]string {
	return map[string]string{
		"id":          "UUID",
		"name":        "Name",
		"created_at":  "CreatedAtInSec",
		"pipeline_id": "PipelineId",
		"status":      "Status",
	}
}

// GetModelName returns table name used as sort field prefix
func (p *PipelineVersion) GetModelName() string {
	return "pipeline_versions."
}