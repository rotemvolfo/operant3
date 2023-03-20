package metrics

import (
	"database/sql"
	"log"
	database "operant/mysql"
)

type Metric struct {
	Name      string
	HighTreshold   int
	LowTreshold int
	Current    int
	UserID string
}

func GetMetricsByUser(userID string) []Metric {
	stmt, err := database.Db.Prepare("select name, high_treshold, low_treshold, current, user_id from Metrics where user_id=?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	metrics, err := ReadRows(rows)
	if err != nil{
		log.Fatal(err)
	}
	return metrics
}

func GetMetricsByUserFilterByTreshold(userID string, highTreshold int, lowTreshold int) []Metric {
	stmt, err := database.Db.Prepare("select name, high_treshold, low_treshold, current, user_id from Metrics WHERE user_id=? AND high_treshold <= ? AND low_treshold >= ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userID, highTreshold, lowTreshold)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	metrics, err := ReadRows(rows)
	if err != nil{
		log.Fatal(err)
	}
	return metrics
}

func ReadRows(rows *sql.Rows) ([]Metric, error) {
	var metrics []Metric
	for rows.Next() {
		var metric Metric
		err := rows.Scan(&metric.Name, &metric.HighTreshold, &metric.LowTreshold, &metric.Current, &metric.UserID)
		if err != nil{
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return metrics, nil;
}