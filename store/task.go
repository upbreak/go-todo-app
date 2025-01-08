package store

import (
	"context"
	"fmt"
	"strconv"

	//"github.com/godror/godror"
	_ "github.com/godror/godror"
	"github.com/upbreak/go-todo-app/entity"
	//"io"
	//"log"
)

func (r *Repository) ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
				t1.IDX,
				t1.TITLE,
-- 				DBMS_LOB.SUBSTR(t1.CONTENT, DBMS_LOB.GETLENGTH(t1.CONTENT)) as content,
				t1.CONTENT,
				t1.REG_DATE
			FROM
				IRIS_NOTICE_BOARD t1
			WHERE
				t1.IS_USE = 'Y'`

	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		fmt.Println("store/task.go ListTasks error")
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) DetailTask(ctx context.Context, db Queryer, idx string) (entity.Task, error) {
	task := entity.Task{}
	sql := `SELECT
				t1.IDX,
				t1.TITLE,
				t1.CONTENT,
				t1.REG_DATE
			FROM
				IRIS_NOTICE_BOARD t1
			WHERE
				t1.IS_USE = 'Y'
			AND t1.IDX = :1`

	idxValue, err := strconv.Atoi(idx)
	if err != nil {
		return task, err
	}
	if err := db.GetContext(ctx, &task, sql, idxValue); err != nil {
		return task, err
	}
	return task, nil
}

func (r *Repository) AddTask(ctx context.Context, db Beginner, t *entity.Task) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("Failed to begin transaction: %v", err)
	}

	sql := `INSERT INTO
				IRIS_NOTICE_BOARD (IDX, SNO, TITLE, CONTENT, SHOW_YN, IS_USE, REG_UNO, REG_USER, REG_DATE)
			VALUES (
					SEQ_IRIS_NOTICE_BOARD.nextval,
					:1,
					:2,
					:3,
					'Y',
					'Y',
					:4,
					:5,
					SYSDATE
			)`
	_, err = tx.ExecContext(ctx, sql, t.SNO, t.TITLE, t.CONTENT, t.REG_UNO, t.REG_USER)

	if err != nil {
		fmt.Println("store/task.go AddTask ExecContext error")
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	//id, err := result.LastInsertId()
	//if err != nil {
	//	tx.Rollback()
	//	fmt.Println("store/task.go AddTask LastInsertId error")
	//	return err
	//}
	//t.IDX = entity.TaskID(id)

	if err := tx.Commit(); err != nil {
		fmt.Println("Failed to commit transaction: %v", err)
	}
	return nil
}
