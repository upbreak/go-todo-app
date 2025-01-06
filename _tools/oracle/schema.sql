CREATE TABLE "users"
(
    "id"       NUMBER(19) NOT NULL,
    "name"     VARCHAR2(20) NOT NULL,
    "password" VARCHAR2(80) NOT NULL,
    "role"     VARCHAR2(80) NOT NULL,
    "created"  DATE NOT NULL,
    "modified" DATE NOT NULL,
    CONSTRAINT PK_USER PRIMARY KEY("id")
);

COMMENT ON TABLE users IS 'go-todo-test 사용자 테이블';
COMMENT ON COLUMN users.id IS '사용자 식별자';
COMMENT ON COLUMN users.name IS '사용자명';
COMMENT ON COLUMN users.password IS '패스워드 해시';
COMMENT ON COLUMN users.role IS '역할';
COMMENT ON COLUMN users.created IS '레코드 작성 시간';
COMMENT ON COLUMN users.modified IS '레코드 수정 시간';

CREATE TABLE "task"
(
    "id"       NUMBER(19) NOT NULL AUTO_INCREMENT COMMENT '테스크 식별자',
--     "user_id"  NUMBER(19) NOT NULL COMMENT '',
    "title"    VARCHAR2(128) NOT NULL COMMENT '테스크 타이틀',
    "status"   VARCHAR2(20)  NOT NULL COMMENT '테스크 상태',
    "created"  DATE NOT NULL COMMENT '테스크 작성 시간',
    "modified" DATE NOT NULL COMMENT '테스크 수정 시간',
    PRIMARY KEY ("id")
)

COMMENT ON TABLE task IS 'go-todo-test task 테이블';
COMMENT ON COLUMN task.id IS '테스크 식별자';
-- COMMENT ON COLUMN task.user_id IS '테스크 식별자';
COMMENT ON COLUMN task.title IS '테스크 타이틀';
COMMENT ON COLUMN task.status IS '테스크 상태';
COMMENT ON COLUMN task.created IS '테스크 작성 시간';
COMMENT ON COLUMN task.modified IS '테스크 수정 시간';