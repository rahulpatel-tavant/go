package tournament

import (
  "fmt"
  "io"
  "bufio"
  "strings"
  "sort"
  "strconv"
  "errors"
)

type Record struct {
  MP int
  W int
  D int
  L int
  P int
  TEAM string
}

type Points []Record
func (a Points) Len() int           { return len(a) }
func (a Points) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Points) Less(i, j int) bool { return a[i].P > a[j].P || (a[i].P == a[j].P && a[i].TEAM < a[j].TEAM) }

func Tally(reader io.Reader, writer io.Writer) error {
  scanner := bufio.NewScanner(reader)
  writer.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
  resultData := make(map[string]map[string]int)
  for scanner.Scan() {
    match_detail := scanner.Text()
    if len(match_detail) > 0 && !strings.HasPrefix(match_detail, "#") {
      match_result, err := TallyMatch(match_detail)
      if err != nil {
        return err
      }
      temp := make(map[string]int)
      temp1 := make(map[string]int)
      if match_result["draw"] == "true" {
        temp["MP"] = 1 + resultData[match_result["winner"]]["MP"]
        temp["W"] = 0 + resultData[match_result["winner"]]["W"]
        temp["D"] = 1 + resultData[match_result["winner"]]["D"]
        temp["L"] = 0 + resultData[match_result["winner"]]["L"]
        temp["P"] = 1 + resultData[match_result["winner"]]["P"]
        resultData[match_result["winner"]] = temp
        temp1["MP"] = 1 + resultData[match_result["loser"]]["MP"]
        temp1["W"] = 0 + resultData[match_result["loser"]]["W"]
        temp1["D"] = 1 + resultData[match_result["loser"]]["D"]
        temp1["L"] = 0 + resultData[match_result["loser"]]["L"]
        temp1["P"] = 1 + resultData[match_result["loser"]]["P"]
        resultData[match_result["loser"]] = temp1
      } else {
        temp["MP"] = 1 + resultData[match_result["winner"]]["MP"]
        temp["W"] = 1 + resultData[match_result["winner"]]["W"]
        temp["D"] = 0 + resultData[match_result["winner"]]["D"]
        temp["L"] = 0 + resultData[match_result["winner"]]["L"]
        temp["P"] = 3 + resultData[match_result["winner"]]["P"]
        resultData[match_result["winner"]] = temp
        temp1["MP"] = 1 + resultData[match_result["loser"]]["MP"]
        temp1["W"] = 0 + resultData[match_result["loser"]]["W"]
        temp1["D"] = 0 + resultData[match_result["loser"]]["D"]
        temp1["L"] = 1 + resultData[match_result["loser"]]["L"]
        temp1["P"] = 0 + resultData[match_result["loser"]]["P"]
        resultData[match_result["loser"]] = temp1
      }
    }
  }
  var records []Record
  for team, data := range resultData {
    records = append(records, Record{TEAM: team, MP: data["MP"], W: data["W"], D: data["D"], L: data["L"], P: data["P"]})
  }
  sort.Sort(Points(records))
  for _, record := range records {
    s := fmt.Sprintf("%-31s|  %-2s|  %-2s|  %-2s|  %-2s|  %s\n", record.TEAM, strconv.Itoa(record.MP), strconv.Itoa(record.W), strconv.Itoa(record.D), strconv.Itoa(record.L), strconv.Itoa(record.P))
    writer.Write([]byte(s))
  }
  return nil
}

func TallyMatch(match_detail string) (map[string]string, error){
  match_details := strings.Split(match_detail, ";")
  if len(match_details) != 3 || (match_details[2] != "win" && match_details[2] != "loss" && match_details[2] != "draw") {
    err := errors.New("Invalid Input")
    return nil, err
  }
  team_1 := match_details[0]
  team_2 := match_details[1]
  result := match_details[2]
  match_result_detail := make(map[string]string)
  if result == "win" {
    match_result_detail["winner"] = team_1
    match_result_detail["loser"] = team_2
  } else if result == "loss"  {
    match_result_detail["winner"] = team_2
    match_result_detail["loser"] = team_1
  } else if result == "draw"  {
    match_result_detail["winner"] = team_1
    match_result_detail["loser"] = team_2
    match_result_detail["draw"] = "true"
  }
  return match_result_detail, nil
}