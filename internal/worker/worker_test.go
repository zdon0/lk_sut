package worker_test

import (
	"strconv"
	"testing"
	"time"

	"lk_sut/internal/testutils"
)

func TestWorkerCommit_no_users(t *testing.T) {
	app := testutils.NewMockApiWithWorker(t, testutils.ExpectGetAllUsers(nil))

	t.Cleanup(app.ClearFunc)

	time.Sleep(testutils.WorkerCommitInterval * 3 / 2)
}

func TestWorkerCommit_already_committed(t *testing.T) {
	expectations := []testutils.Opt{
		testutils.ExpectGetAllUsers(map[string]string{
			testutils.Login: testutils.Password,
		}),
		testutils.ExpectGetUserLastLogin(testutils.Login, strconv.FormatInt(time.Now().UnixNano(), 10)),
	}

	app := testutils.NewMockApiWithWorker(t, expectations...)

	t.Cleanup(app.ClearFunc)

	time.Sleep(testutils.WorkerCommitInterval * 3 / 2)
}

func TestWorkerCommit(t *testing.T) {
	// TODO: Schedule and Lesson commit Handlers
	t.SkipNow()

	expectations := []testutils.Opt{
		testutils.ExpectGetAllUsers(map[string]string{
			testutils.Login: testutils.Password,
		}),
		testutils.ExpectGetUserLastLogin(testutils.Login, ""),

		testutils.RegisterUserInitHandler(),
		testutils.RegisterUserAuthHandler(testutils.Login, testutils.Password),
		//testutils.RegisterScheduleHandler(),
		//testutils.RegisterLessonCommitHandler(),

		testutils.ExpectSetUserLastLogin(testutils.Login),
	}

	app := testutils.NewMockApiWithWorker(t, expectations...)

	t.Cleanup(app.ClearFunc)

	time.Sleep(testutils.WorkerCommitInterval * 3 / 2)
}
