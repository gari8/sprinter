// DO NOT EDIT.

package tmpl

import "text/template"

var layout = "-- Makefile --\ninit:\n\tgo mod init @@.ImportPath@@ && go mod tidy && make build && make up_db && make up\nbuild:\n\tdocker-compose -f deployments/docker-compose.yml build --no-cache\nup:\n\tdocker-compose -f deployments/docker-compose.yml up api\nup_db:\n\tdocker-compose -f deployments/docker-compose.yml up -d db\nrun_api:\n\tdocker-compose -f deployments/docker-compose.yml run api ash\nrun_db:\n\tdocker-compose -f deployments/docker-compose.yml run db ash\nexec_api:\n\tdocker-compose -f deployments/docker-compose.yml exec api ash\nexec_db:\n\tdocker-compose -f deployments/docker-compose.yml exec db ash\n-- README.md --\n# @@.ImportPath@@\n\n- this app was made by github.com/gari8/sprinter\n\n- read MakeFile\n-- build/Dockerfile --\nFROM golang:alpine\n\nENV GO113MODULE=on\n\nWORKDIR /app\n\nCOPY .. .\n\nRUN go mod tidy\n\nRUN apk update && apk add --no-cache git && go get github.com/cespare/reflex\n-- deployments/docker-compose.yml --\nversion: \"3.8\"\nservices:\n  api:\n    container_name: api\n    build:\n      context: ../\n      dockerfile: ../build/Dockerfile\n    tty: true\n    restart: always\n    volumes:\n      - ..:/app\n    ports:\n      - 8080:8080\n    environment:\n      PORT: 8080\n      @@ if eq .DataBase \"MySQL\" -@@\n      DRIVER: \"mysql\"\n      DATABASE_URL: \"mysql:mysql@tcp(db:3306)/mysql?charset=utf8&parseTime=true\"\n      @@ else @@\n      DRIVER: \"postgres\"\n      DATABASE_URL: \"host=db user=postgres dbname=postgres password=postgres sslmode=disable\"\n      @@ end @@\n    command: sh scripts/build-local.sh\n  db:\n    @@ if eq .DataBase \"MySQL\" -@@\n    image: mysql:alpine\n    container_name: db\n    environment:\n      MYSQL_ROOT_PASSWORD: mysql\n      MYSQL_DATABASE: mysql\n      MYSQL_USER: mysql\n      MYSQL_PASSWORD: mysql\n      TZ: 'Asia/Tokyo'\n    command: mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci\n    volumes:\n      - ../test/database/init:/docker-entrypoint-initdb.d\n    ports:\n      - 3306:3306\n    @@ else @@\n    image: postgres:10-alpine\n    container_name: db\n    ports:\n      - 5432:5432\n    environment:\n      POSTGRES_USER: postgres\n      POSTGRES_PASSWORD: postgres\n      PGPASSWORD: postgres\n      POSTGRES_DB: postgres\n      DATABASE_HOST: localhost\n      TZ: 'Asia/Tokyo'\n    command: postgres -c log_destination=stderr -c log_statement=all -c log_connections=on -c log_disconnections=on\n    logging:\n      options:\n        max-size: \"10k\"\n        max-file: \"5\"\n    volumes:\n      - ../test/database/init:/docker-entrypoint-initdb.d\n    @@ end @@\n  \n-- scripts/build-local.sh --\n#!/bin/bash\n\ngo mod tidy\nreflex -r '(\\.go$|go\\.mod)' -s go run cmd/main.go\n-- test/database/init/1_init.sql --\n\nDROP TABLE IF EXISTS samples;\n\nCREATE TABLE IF NOT EXISTS samples\n(\n   id SERIAL NOT NULL,\n   text TEXT NOT NULL,\n   PRIMARY KEY (id)\n);\n\nINSERT INTO samples(text) VALUES ('sample');\n\n"
var OnionTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage cmd\n-- internal/@@.ImportPath@@/infrastructure/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n"))
var MVCTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage main\n\nimport (\n\t\"fmt\"\n\t\"net/http\"\n\t\"os\"\n\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/config/database/conf\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/controller\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/model\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/router\"\n)\n\nconst defaultPort = \"8080\"\n\nfunc main() {\n\tport := os.Getenv(\"PORT\")\n\tif port == \"\" {\n\t\tport = defaultPort\n\t}\n\n\tconn, err := conf.NewDatabaseConnection()\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\tif conn == nil {\n\t\tpanic(err)\n\t}\n\tdefer func() {\n\t\tif conn != nil {\n\t\t\tif err := conn.Close(); err != nil {\n\t\t\t\tpanic(err)\n\t\t\t}\n\t\t}\n\t}()\n\n\tfmt.Println(`\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * \n  *        ####    #####    #####     ####    ##  ##   ######   ######   #####  *\n  *      ##  ##   ##  ##   ##  ##     ##     ### ##     ##     ##       ##  ##  *\n  *     ##       ##  ##   ##  ##     ##     ######     ##     ##       ##  ##   *\n  *     ####    #####    #####      ##     ######     ##     ####     #####     *\n  *       ##   ##       ####       ##     ## ###     ##     ##       ####       *\n  *  ##  ##   ##       ## ##      ##     ##  ##     ##     ##       ## ##       *\n  *  ####    ##       ##  ##    ####    ##  ##     ##     ######   ##  ##       *\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n\n\tHTML:\tGET http://localhost:8080\n    API:\tGET http://localhost:8080/api/v1\n    `)\n\n\tm := model.NewModel(conn)\n\tc := controller.NewController(m)\n\ts := router.NewRouter()\n\ts.Router(c)\n\n\t_ = http.ListenAndServe(\":\"+port, s.Route)\n}\n-- internal/@@.ImportPath@@/config/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n-- internal/@@.ImportPath@@/controller/controller_base.go --\npackage controller\n\nimport (\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/model\"\n)\n\ntype Controller struct {\n\tSampleController\n}\n\nfunc NewController(model model.Model) Controller {\n\tc := Controller{}\n\tc.SampleController = NewSampleController(model)\n\treturn c\n}\n-- internal/@@.ImportPath@@/controller/controller_util.go --\npackage controller\n\nimport (\n\t\"fmt\"\n\t\"html/template\"\n\t\"io/ioutil\"\n\t\"os\"\n\t\"path/filepath\"\n\t\"strings\"\n)\n\ntype response struct {\n\tStatus int\n\tResult interface{}\n}\n\nfunc parseTemplate(dir string, fileName string) (*template.Template, error) {\n\ttmpl := template.New(\"\")\n\n\tvar layout string\n\n\tif err := filepath.Walk(\"internal/mymvc/views/layout\", func(path string, info os.FileInfo, err error) error {\n\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\tif !info.IsDir() && (strings.HasSuffix(path, \".html\") || strings.HasSuffix(path, \".js\")) {\n\t\t\tfile, err := ioutil.ReadFile(path)\n\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\n\t\t\tlayout += string(file)\n\t\t}\n\n\t\treturn nil\n\t}); err != nil {\n\t\tfmt.Println(err)\n\t}\n\n\tif err := filepath.Walk(\"internal/mymvc/views/\"+dir, func(path string, info os.FileInfo, err error) error {\n\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\tif !info.IsDir() && (strings.HasSuffix(path, \".html\") || strings.HasSuffix(path, \".js\")) {\n\t\t\tfile, err := ioutil.ReadFile(path)\n\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\n\t\t\tfilename := strings.Replace(path, \"views/\"+dir, \"\", -1)\n\n\t\t\tif strings.Contains(filename, fileName) {\n\t\t\t\ttmpl = tmpl.New(filename)\n\n\t\t\t\ttmpl, err = tmpl.Parse(string(file) + layout)\n\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn err\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\n\t\treturn nil\n\t}); err != nil {\n\t\treturn nil, err\n\t}\n\n\treturn tmpl, nil\n}\n-- internal/@@.ImportPath@@/controller/sample_controller.go --\npackage controller\n\nimport (\n\t\"encoding/json\"\n\t\"log\"\n\t\"net/http\"\n\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/model\"\n)\n\ntype SampleController struct {\n\tmodel.Model\n}\n\nfunc NewSampleController(sm model.Model) SampleController {\n\treturn SampleController{sm}\n}\n\nfunc (s SampleController) SampleIndex(w http.ResponseWriter, r *http.Request) {\n\tsamples, err := s.SampleModel.Fetch()\n\n\tif err != nil {\n\t\thttp.Error(w, err.Error(), http.StatusNotFound)\n\t}\n\n\tresp := &response{\n\t\tStatus: http.StatusOK,\n\t\tResult: samples,\n\t}\n\n\tres, err := json.Marshal(resp)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\t_, _ = w.Write(res)\n}\n\nfunc (s SampleController) SampleHTML(w http.ResponseWriter, r *http.Request) {\n\ttmpl, err := parseTemplate(\"sample\", \"index\")\n\n\tif err != nil {\n\t\tlog.Fatal(\"err :\", err)\n\t}\n\n\tif err := tmpl.Execute(w, nil); err != nil {\n\t\tlog.Printf(\"failed to execute template: %v\", err)\n\t}\n}\n-- internal/@@.ImportPath@@/model/model_base.go --\npackage model\n\nimport \"database/sql\"\n\ntype Model struct {\n\tSampleModel\n}\n\nfunc NewModel(conn *sql.DB) Model {\n\tm := Model{}\n\tm.SampleModel = NewSampleModel(conn)\n\treturn m\n}\n-- internal/@@.ImportPath@@/model/sample_model.go --\npackage model\n\nimport \"database/sql\"\n\ntype SampleModel struct {\n\tconn *sql.DB\n}\n\nfunc NewSampleModel(conn *sql.DB) SampleModel {\n\treturn SampleModel{conn}\n}\n\ntype Sample struct {\n\tID   int64 `json:\"id\" db:\"id\"`\n\tText string `json:\"text\" db:\"text\"`\n}\n\nfunc (s *SampleModel) Fetch() ([]*Sample, error) {\n\tvar samples []*Sample\n\trows, err := s.conn.Query(\"SELECT id, text FROM samples;\")\n\tif rows == nil {\n\t\treturn nil, err\n\t}\n\tfor rows.Next() {\n\t\tsample := &Sample{}\n\t\terr = rows.Scan(&sample.ID, &sample.Text)\n\t\tif err == nil {\n\t\t\tsamples = append(samples, sample)\n\t\t}\n\t}\n\treturn samples, err\n}\n-- internal/@@.ImportPath@@/router/router.go --\npackage router\n\nimport (\n\t\"github.com/go-chi/chi\"\n\t\"github.com/go-chi/chi/middleware\"\n\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/controller\"\n)\n\ntype Server struct {\n\tRoute *chi.Mux\n}\n\nfunc NewRouter() *Server {\n\treturn &Server{\n\t\tRoute: chi.NewRouter(),\n\t}\n}\n\nfunc (s *Server) Router(c controller.Controller) {\n\ts.Route.Use(middleware.Logger)\n\ts.Route.Use(middleware.Recoverer)\n\ts.Route.Route(\"/\", func(r chi.Router) {\n\t\tr.Get(\"/\", c.SampleController.SampleHTML)\n\t})\n\ts.Route.Route(\"/api/v1\", func(r chi.Router) {\n\t\tr.Get(\"/\", c.SampleController.SampleIndex)\n\t\t// TODO\n\t})\n}\n-- internal/@@.ImportPath@@/views/layout/_footer.html --\n{{ define \"footer\" }}\n    <footer>sample</footer>\n    </body>\n    </html>\n{{ end }}\n-- internal/@@.ImportPath@@/views/layout/_header.html --\n{{ define \"header\" }}\n    <!DOCTYPE html>\n    <html lang=\"en\">\n    <head>\n        <meta charset=\"UTF-8\">\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n        <title>Sample</title>\n        <style>\n            * {\n                margin: 0;\n                padding: 0;\n                box-sizing: border-box;\n            }\n        </style>\n    </head>\n    <body>\n    <header><a href=\"/\">sample</a></header>\n{{ end }}\n-- internal/@@.ImportPath@@/views/sample/index.html --\n{{ template \"header\" }}\n<h1><i>SPRINTER</i></h1>\n{{ template \"footer\" }}\n"))
var MinimumTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage main\n\nimport (\n\t\"fmt\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/config/database/conf\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/router\"\n\t\"net/http\"\n\t\"os\"\n)\n\nconst defaultPort = \"8080\"\n\nfunc main() {\n\tport := os.Getenv(\"PORT\")\n\tif port == \"\" {\n\t\tport = defaultPort\n\t}\n\n\tconn, err := conf.NewDatabaseConnection()\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\tif conn == nil {\n\t\tpanic(err)\n\t}\n\tdefer func() {\n\t\tif conn != nil {\n\t\t\tif err := conn.Close(); err != nil {\n\t\t\t\tpanic(err)\n\t\t\t}\n\t\t}\n\t}()\n\n\tfmt.Println(`\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * \n  *        ####    #####    #####     ####    ##  ##   ######   ######   #####  *\n  *      ##  ##   ##  ##   ##  ##     ##     ### ##     ##     ##       ##  ##  *\n  *     ##       ##  ##   ##  ##     ##     ######     ##     ##       ##  ##   *\n  *     ####    #####    #####      ##     ######     ##     ####     #####     *\n  *       ##   ##       ####       ##     ## ###     ##     ##       ####       *\n  *  ##  ##   ##       ## ##      ##     ##  ##     ##     ##       ## ##       *\n  *  ####    ##       ##  ##    ####    ##  ##     ##     ######   ##  ##       *\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n\n    API:\tGET http://localhost:8080/api/v1\n    `)\n\n\ts := router.NewRouter(conn)\n\ts.Router()\n\n\t_ = http.ListenAndServe(\":\"+port, s.Route)\n}\n-- internal/@@.ImportPath@@/config/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n-- internal/@@.ImportPath@@/router/router.go --\npackage router\n\nimport (\n\t\"database/sql\"\n\t\"encoding/json\"\n\t\"github.com/go-chi/chi\"\n\t\"github.com/go-chi/chi/middleware\"\n\t\"log\"\n\t\"net/http\"\n)\n\ntype Server struct {\n\tRoute *chi.Mux\n\tConn *sql.DB\n}\n\nfunc NewRouter(conn *sql.DB) *Server {\n\treturn &Server{\n\t\tRoute: chi.NewRouter(),\n\t\tConn: conn,\n\t}\n}\n\ntype Sample struct {\n\tID   int64 `json:\"id\" db:\"id\"`\n\tText string `json:\"text\" db:\"text\"`\n}\n\nfunc (s *Server) Router() {\n\ts.Route.Use(middleware.Logger)\n\ts.Route.Use(middleware.Recoverer)\n\ts.Route.Route(\"/api/v1\", func(r chi.Router) {\n\t\tr.Get(\"/\", func(w http.ResponseWriter, r *http.Request) {\n\t\t\tvar samples []*Sample\n\t\t\trows, err := s.Conn.Query(\"SELECT id, text FROM samples;\")\n\t\t\tif rows == nil {\n\t\t\t\tlog.Fatalln(err)\n\t\t\t}\n\t\t\tfor rows.Next() {\n\t\t\t\tsample := &Sample{}\n\t\t\t\terr = rows.Scan(&sample.ID, &sample.Text)\n\t\t\t\tif err == nil {\n\t\t\t\t\tsamples = append(samples, sample)\n\t\t\t\t}\n\t\t\t}\n\n\t\t\tres, err := json.Marshal(samples)\n\t\t\tif err != nil {\n\t\t\t\tlog.Fatal(err)\n\t\t\t}\n\n\t\t\t_, _ = w.Write(res)\n\t\t})\n\t})\n}\n"))
var CleanTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage main\n\nimport (\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/infrastructure/database/conf\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/infrastructure/server\"\n)\n\nfunc main() {\n\tconn, err := conf.NewDatabaseConnection()\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\ts := server.NewServer(conn)\n\ts.Serve()\n}\n-- internal/@@.ImportPath@@/domain/sample.go --\npackage domain\n\ntype Sample struct {\n\tID   uint32 `json:\"id\"`\n\tText string `json:\"text,omitempty\"`\n}\n-- internal/@@.ImportPath@@/infrastructure/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n-- internal/@@.ImportPath@@/infrastructure/server/sample_router.go --\npackage server\n\nimport (\n\t\"fmt\"\n\t\"github.com/gari8/sprinter\"\n\t\"github.com/go-chi/chi\"\n\tsampleController \"@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/controllers/sample\"\n\tsamplePresenter \"@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/presenters/sample\"\n\tsampleRepository \"@@.ImportPath@@/internal/@@.ImportPath@@/interfaces/repositories/sample\"\n\tsampleService \"@@.ImportPath@@/internal/@@.ImportPath@@/usecases/services/sample\"\n\tsampleUseCase \"@@.ImportPath@@/internal/@@.ImportPath@@/usecases/usecases/sample\"\n\t\"net/http\"\n)\n\nfunc sampleRouter(s Server) http.Handler {\n\tr := chi.NewRouter()\n\n\trepo := sampleRepository.NewSampleRepository(s.Conn)\n\tservice := sampleService.NewSampleService(repo)\n\tpresenter := samplePresenter.NewGetSamplePresenter()\n\tinteractor := sampleUseCase.NewGetSampleInteractor(presenter, service)\n\tcontroller := sampleController.NewSampleController(interactor)\n\tr.Get(\"/\", sprinter.Handle(controller.GetSamples))\n\tr.Post(\"/\", func(writer http.ResponseWriter, request *http.Request) {\n\t\tfmt.Println(\"=======hi\")\n\t})\n\treturn r\n}\n-- internal/@@.ImportPath@@/infrastructure/server/server.go --\npackage server\n\nimport (\n\t\"database/sql\"\n\t\"github.com/go-chi/chi\"\n\t\"github.com/go-chi/chi/middleware\"\n\t\"net/http\"\n)\n\ntype Server struct {\n\tRoute *chi.Mux\n\tConn *sql.DB\n}\n\nfunc NewServer(conn *sql.DB) Server {\n\treturn Server{\n\t\tRoute: chi.NewRouter(),\n\t\tConn: conn,\n\t}\n}\n\nfunc (s Server) Serve() {\n\ts.Route.Use(middleware.Logger)\n\ts.Route.Use(middleware.Recoverer)\n\ts.Route.Route(\"/api/v1\", func(r chi.Router) {\n\t\tr.Mount(\"/sample\", sampleRouter(s))\n\t})\n\terr := http.ListenAndServe(\":8080\", s.Route)\n\tif err != nil {\n\t\tpanic(err)\n\t}\n}\n-- internal/@@.ImportPath@@/interfaces/controllers/sample/sample_controller.go --\npackage sample\n\nimport (\n\t\"context\"\n\t\"github.com/gari8/sprinter\"\n\tsampleUseCase \"@@.ImportPath@@/internal/@@.ImportPath@@/usecases/usecases/sample\"\n)\n\ntype SampleController interface {\n\tGetSamples(ctx context.Context) sprinter.Response\n}\n\ntype sampleController struct {\n\tsampleUseCase.GetSampleInputPort\n}\n\nfunc NewSampleController(\n\tinputPort sampleUseCase.GetSampleInputPort,\n\t) SampleController {\n\treturn sampleController{inputPort}\n}\n\nfunc (c sampleController) GetSamples(ctx context.Context) sprinter.Response {\n\treturn c.GetSampleInputPort.GetSamples()\n}\n-- internal/@@.ImportPath@@/interfaces/presenters/sample/get_sample_presenter.go --\npackage sample\n\nimport (\n\t\"github.com/gari8/sprinter\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/domain\"\n\t\"net/http\"\n)\n\ntype GetSamplePresenter struct {\n\n}\n\nfunc NewGetSamplePresenter() GetSamplePresenter {\n\treturn GetSamplePresenter{}\n}\n\nfunc (p GetSamplePresenter) CreateResponse(samples []domain.Sample) sprinter.Response {\n\treturn sprinter.Response{\n\t\tCode: http.StatusOK,\n\t\tObject: samples,\n\t}\n}\n-- internal/@@.ImportPath@@/interfaces/repositories/sample/sample_repository.go --\npackage sample\n\nimport (\n\t\"database/sql\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/domain\"\n)\n\ntype SampleRepository struct {\n\tconn *sql.DB\n}\n\nfunc NewSampleRepository(conn *sql.DB) SampleRepository {\n\treturn SampleRepository{conn}\n}\n\nfunc (s SampleRepository) GetSamples() ([]domain.Sample, error) {\n\tvar samples []domain.Sample\n\trows, err := s.conn.Query(\"SELECT id, text FROM samples;\")\n\tif rows == nil {\n\t\treturn nil, err\n\t}\n\tfor rows.Next() {\n\t\tvar sample domain.Sample\n\t\terr = rows.Scan(&sample.ID, &sample.Text)\n\t\tif err == nil {\n\t\t\tsamples = append(samples, sample)\n\t\t}\n\t}\n\treturn samples, err\n}\n-- internal/@@.ImportPath@@/usecases/services/sample/sample_service.go --\npackage sample\n\nimport (\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/domain\"\n)\n\ntype Repository interface {\n\tGetSamples() ([]domain.Sample, error)\n}\n\ntype SampleService struct {\n\trepository Repository\n}\n\nfunc NewSampleService(repository Repository) SampleService {\n\treturn SampleService{repository: repository}\n}\n\nfunc (s SampleService) GetSamples() ([]domain.Sample, error) {\n\treturn s.repository.GetSamples()\n}\n-- internal/@@.ImportPath@@/usecases/usecases/sample/get_sample_input_port.go --\npackage sample\n\nimport (\n\t\"github.com/gari8/sprinter\"\n)\n\ntype GetSampleInputPort interface {\n\tGetSamples () sprinter.Response\n}\n-- internal/@@.ImportPath@@/usecases/usecases/sample/get_sample_interactor.go --\npackage sample\n\nimport (\n\t\"github.com/gari8/sprinter\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/usecases/services/sample\"\n\t\"net/http\"\n)\n\ntype GetSampleInteractor struct {\n\toutputPort GetSampleOutputPort\n\tsampleService sample.SampleService\n}\n\nfunc NewGetSampleInteractor(\n\toutputPort GetSampleOutputPort,\n\tsampleService sample.SampleService,\n\t) GetSampleInteractor {\n\treturn GetSampleInteractor{\n\t\toutputPort: outputPort,\n\t\tsampleService: sampleService,\n\t}\n}\n\nfunc (i GetSampleInteractor) GetSamples() sprinter.Response {\n\tsamples, err := i.sampleService.GetSamples()\n\tif err != nil {\n\t\treturn sprinter.Response{\n\t\t\tCode: http.StatusInternalServerError,\n\t\t\tObject: err,\n\t\t}\n\t}\n\treturn i.outputPort.CreateResponse(samples)\n}\n-- internal/@@.ImportPath@@/usecases/usecases/sample/get_sample_output_port.go --\npackage sample\n\nimport (\n\t\"github.com/gari8/sprinter\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/domain\"\n)\n\ntype GetSampleOutputPort interface {\n\tCreateResponse(samples []domain.Sample) sprinter.Response\n}\n"))
