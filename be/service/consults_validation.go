package service

import (
	"consult-scheduler/model"
	"fmt"
	"sort"
	"time"
)

func validateConsultTimes(
	consults []model.DayTemplateConsultDto,
	consultTypes map[uint]model.ConsultType,
) (bool, error) {
	sortedConsults := make([]model.DayTemplateConsultDto, len(consults))
	copy(sortedConsults, consults)

	// sort consults asc
	sort.Slice(sortedConsults, func(i, j int) bool {
		return sortedConsults[i].Time.GetTime().Before(sortedConsults[j].Time.GetTime())
	})

	for i := range sortedConsults {
		// no need to check for overlapping time on the last consult
		if i == len(sortedConsults)-1 {
			break
		}

		consult := sortedConsults[i]
		consultType := consultTypes[consult.ConsultTypeID]
		consultTime := consult.Time.GetTime()
		// no handling of if time overflows through midnight, because there should be no consults at these ungodly hours :)
		consultEndTime := consultTime.Add(time.Duration(consultType.DurationMinutes) * time.Minute)

		nextConsult := sortedConsults[i+1]
		nextConsultTime := nextConsult.Time.GetTime()
		if nextConsultTime.Before(consultEndTime) {
			return false, model.CustomValidationError{
				FieldName: "time",
				Reason:    fmt.Sprintf("Overlapping times %v and %v", consultTime, nextConsultTime),
			}
		}

	}

	return true, nil
}
