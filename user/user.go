package user

import (
	"fmt"

	"go-utils/user/snowflake"
)

var ids = make([]uint64, 0)

// GenerateUserIds 一次生成十个
func generateUserIds() {
	fmt.Println("generateUserIds start...")
	machineNode := 1
	node, err := snowflake.NewNode(int64(machineNode))
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		ids = append(ids, uint64(node.Generate().Int64()))
	}
}

// GetUserID 从切片获取
func GetUserID() (id uint64) {
	if len(ids) < 1 {
		generateUserIds()
	}
	id = ids[0]
	// 取出后从切片删除
	if len(ids) == 1 {
		ids = ids[:0]
	} else {
		ids = ids[1:]
	}
	return
}
