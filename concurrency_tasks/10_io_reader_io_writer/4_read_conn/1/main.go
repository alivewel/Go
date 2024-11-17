package main

import (
	"fmt"
	"io"
	"strings"
)

func (c *client) GetAccrualsByEmployeeId(ctx context.Context, employeeId int64, limit int64, offset int64, dateFrom string, dateTo string) (accruals []Accrual, total int64, err error) {
	urlQuery := fmt.Sprintf("%s/api/v1/accruals?employee_id=%d&limit=%d&offset=%d&dt_from=%s&dt_to=%s", c.config.Url, employeeId, limit, offset, dateFrom, dateTo)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, urlQuery, http.NoBody)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
	resp, err := c.client.Do(req)
	if err != nil {
		err = fmt.Errorf("ошибка запроса GetAccrualsByEmployeeId :%s\nError: %s\nEmployeeID: %d", urlQuery, err, employeeId)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("ошибка чтения тела ответа GetAccrualsByEmployeeId :%s\nError: %s\nEmployeeID: %d", urlQuery, err, employeeId)
		return
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		err = fmt.Errorf("ошибка статуса запроса GetAccrualsByEmployeeId :%s\nError: %s\nEmployeeID: %d\nStatusCode: %d\nRawBody: %s", urlQuery, err, employeeId, resp.StatusCode, rawBody)
		return
	}

	var r AccrualsResponse
	err = json.Unmarshal(rawBody, &r)
	if err != nil {
		err = fmt.Errorf("ошибка разбора тела ответа GetAccrualsByEmployeeId :%s\nError: %s\nEmployeeID: %d", urlQuery, err, employeeId)
		return
	}

	accruals = r.Data.Items
	total = int64(r.Data.Range.Count)
	return
}

func main() {
	r := strings.NewReader("some string: hello world")

	// read 1024 bytes at a time
	buf := make([]byte, 10)

	// Make a buffer to store the data
	for {
		// Read from file into buffer
		n, err := r.Read(buf)

		// If bytes were read (n > 0), process them
		if n > 0 {
			fmt.Println("Received", n, "bytes:", string(buf[:n]))
		}

		// Check for errors, but handle EOF properly
		if err != nil {
			if err == io.EOF {
				break // End of file, stop reading
			}
			panic(err) // Handle other potential errors
		}
	}
}
