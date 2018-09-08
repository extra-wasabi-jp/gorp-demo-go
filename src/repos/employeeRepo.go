package repos
import (
    "fmt"
    "../entity"
)

func (v Repo) GetEmployeeAll() []entity.Employee {
    var list [] entity.Employee
    _, err := dbmap.Select(&list, "select * from employee where is_actived = '1';")
    if err != nil {
        fmt.Printf("GetEmployeeAll error")
        panic(err)
    }

    return list
}

func (v Repo) GetEmployeeByLoginId(loginId string) entity.Employee {
    var employee entity.Employee
    err := dbmap.SelectOne(&employee, "select * from employee where login_id = ?", loginId)
    if err != nil {
        fmt.Printf("GetEmployeeByLoginId error")
        panic(err)
    }
    return employee
}
