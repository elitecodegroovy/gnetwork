package sqlstore

import (
	"fmt"
	"github.com/elitecodegroovy/gnetwork/pkg/models"
)

func GetOrgUsersForTest(query *models.GetOrgUsersQuery) error {
	query.Result = make([]*models.OrgUserDTO, 0)
	sess := x.Table("org_user")
	sess.Join("LEFT ", x.Dialect().Quote("user"), fmt.Sprintf("org_user.user_id=%s.id", x.Dialect().Quote("user")))
	sess.Where("org_user.org_id=?", query.OrgId)
	sess.Cols("org_user.org_id", "org_user.user_id", "user.email", "user.login", "org_user.role")

	err := sess.Find(&query.Result)
	return err
}
