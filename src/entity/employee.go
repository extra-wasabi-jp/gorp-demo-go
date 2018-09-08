package entity
import (
    "time"
    "database/sql"
)

type Employee struct {
    EmployeeId        int64          `db:"employee_id"`
    LoginId           string         `db:"login_id"`
    EmpNo             string         `db:"emp_no"`
    Email             string         `db:"email"`
    Password          string         `db:"password"`
    PasswordLsgChgDt  string         `db:"password_lst_chg_dt"`
    UserNmJa          string         `db:"user_nm_ja"`
    UserNmKa          string         `db:"user_nm_ka"`
    UserNmEn          string         `db:"user_nm_en"`
    JoinDt            string         `db:"join_dt"`
    QuitDt            sql.NullString `db:"quit_dt"`
    VersionNo         int8           `db:"version_no"`
    CreatedAt         time.Time      `db:"created_at"`
    CreatedBy         string         `db:"created_by"`
    UpdatedAt         time.Time      `db:"updated_at"`
    UpdatedBy         string         `db:"updated_by"`
    IsActived         string         `db:"is_actived"`
}

