// DO NOT EDIT.

package tmpl

import "text/template"

var layout = "-- Makefile --\ninit:\n\tgo mod init @@.ImportPath@@ && go mod tidy && make build && make up_db && make up\nbuild:\n\tdocker-compose -f deployments/docker-compose.yml build --no-cache\nup:\n\tdocker-compose -f deployments/docker-compose.yml up api\nup_db:\n\tdocker-compose -f deployments/docker-compose.yml up -d db\nrun_api:\n\tdocker-compose -f deployments/docker-compose.yml run api ash\nrun_db:\n\tdocker-compose -f deployments/docker-compose.yml run db ash\n-- README.md --\n# @@.ImportPath@@\n\n- this app was made by github.com/gari8/sprinter\n\n- read MakeFile\n-- build/Dockerfile --\nFROM golang:alpine\n\nENV GO113MODULE=on\n\nWORKDIR /app\n\nCOPY ../@@.ImportPath@@ .\n\nRUN go mod tidy\n\nRUN apk update && apk add --no-cache git && go get github.com/cespare/reflex\n-- deployments/docker-compose.yml --\nversion: \"3.8\"\nservices:\n  api:\n    container_name: api\n    build:\n      context: ../@@.ImportPath@@\n      dockerfile: ../build/Dockerfile\n    tty: true\n    restart: always\n    volumes:\n      - ..:/app\n    ports:\n      - 8080:8080\n    environment:\n      PORT: 8080\n      @@ if eq .DataBase \"MySQL\" -@@\n      DRIVER: \"mysql\"\n      DATABASE_URL: \"mysql:mysql@tcp(@@.ImportPath@@db:3306)/mysql?charset=utf8&parseTime=true\"\n      @@ else @@\n      DRIVER: \"postgres\"\n      DATABASE_URL: \"host=@@.ImportPath@@db user=postgres dbname=postgres password=postgres sslmode=disable\"\n      @@ end @@\n    command: sh scripts/build-local.sh\n  db:\n    @@ if eq .DataBase \"MySQL\" -@@\n    image: mysql:alpine\n    container_name: db\n    environment:\n      MYSQL_ROOT_PASSWORD: mysql\n      MYSQL_DATABASE: mysql\n      MYSQL_USER: mysql\n      MYSQL_PASSWORD: mysql\n      TZ: 'Asia/Tokyo'\n    command: mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci\n    volumes:\n      - ../test/database/init:/docker-entrypoint-initdb.d\n    ports:\n      - 3306:3306\n    @@ else @@\n    image: postgres:10-alpine\n    container_name: db\n    ports:\n      - 5432:5432\n    environment:\n      POSTGRES_USER: postgres\n      POSTGRES_PASSWORD: postgres\n      PGPASSWORD: postgres\n      POSTGRES_DB: postgres\n      DATABASE_HOST: localhost\n      TZ: 'Asia/Tokyo'\n    command: postgres -c log_destination=stderr -c log_statement=all -c log_connections=on -c log_disconnections=on\n    logging:\n      options:\n        max-size: \"10k\"\n        max-file: \"5\"\n    volumes:\n      - ../test/database/init:/docker-entrypoint-initdb.d\n    @@ end @@\n  \n-- scripts/build-local.sh --\n#!/bin/bash\n\ngo mod tidy\nreflex -r '(\\.go$|go\\.mod)' -s go run cmd/main.go\n-- test/database/init/1_init.sql --\n\nDROP TABLE IF EXISTS samples;\n\nCREATE TABLE IF NOT EXISTS samples\n(\n   id SERIAL NOT NULL,\n   text TEXT NOT NULL,\n   PRIMARY KEY (id)\n);\n\nINSERT INTO samples(text) VALUES ('sample');\n\n"
var OnionTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage cmd\n-- internal/@@.ImportPath@@/infrastructure/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n"))
var MVCTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage main\n\nimport (\n\t\"fmt\"\n\t\"net/http\"\n\t\"os\"\n\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/config/database/conf\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/controller\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/model\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/router\"\n)\n\nconst defaultPort = \"8080\"\n\nfunc main() {\n\tport := os.Getenv(\"PORT\")\n\tif port == \"\" {\n\t\tport = defaultPort\n\t}\n\n\tconn, err := conf.NewDatabaseConnection()\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\tif conn == nil {\n\t\tpanic(err)\n\t}\n\tdefer func() {\n\t\tif conn != nil {\n\t\t\tif err := conn.Close(); err != nil {\n\t\t\t\tpanic(err)\n\t\t\t}\n\t\t}\n\t}()\n\n\tfmt.Println(`\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * \n  *        ####    #####    #####     ####    ##  ##   ######   ######   #####  *\n  *      ##  ##   ##  ##   ##  ##     ##     ### ##     ##     ##       ##  ##  *\n  *     ##       ##  ##   ##  ##     ##     ######     ##     ##       ##  ##   *\n  *     ####    #####    #####      ##     ######     ##     ####     #####     *\n  *       ##   ##       ####       ##     ## ###     ##     ##       ####       *\n  *  ##  ##   ##       ## ##      ##     ##  ##     ##     ##       ## ##       *\n  *  ####    ##       ##  ##    ####    ##  ##     ##     ######   ##  ##       *\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n\n\tHTML:\tGET http://localhost:8080\n    API:\tGET http://localhost:8080/api/v1\n    `)\n\n\tm := model.NewModel(conn)\n\tc := controller.NewController(m)\n\ts := router.NewRouter()\n\ts.Router(c)\n\n\t_ = http.ListenAndServe(\":\"+port, s.Route)\n}\n-- internal/@@.ImportPath@@/config/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n-- internal/@@.ImportPath@@/controller/controller_base.go --\npackage controller\n\nimport (\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/model\"\n)\n\ntype Controller struct {\n\tSampleController\n}\n\nfunc NewController(model model.Model) Controller {\n\tc := Controller{}\n\tc.SampleController = NewSampleController(model)\n\treturn c\n}\n-- internal/@@.ImportPath@@/controller/controller_util.go --\npackage controller\n\nimport (\n\t\"fmt\"\n\t\"html/template\"\n\t\"io/ioutil\"\n\t\"os\"\n\t\"path/filepath\"\n\t\"strings\"\n)\n\ntype response struct {\n\tStatus int\n\tResult interface{}\n}\n\nfunc parseTemplate(dir string, fileName string) (*template.Template, error) {\n\ttmpl := template.New(\"\")\n\n\tvar layout string\n\n\tif err := filepath.Walk(\"internal/mymvc/views/layout\", func(path string, info os.FileInfo, err error) error {\n\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\tif !info.IsDir() && (strings.HasSuffix(path, \".html\") || strings.HasSuffix(path, \".js\")) {\n\t\t\tfile, err := ioutil.ReadFile(path)\n\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\n\t\t\tlayout += string(file)\n\t\t}\n\n\t\treturn nil\n\t}); err != nil {\n\t\tfmt.Println(err)\n\t}\n\n\tif err := filepath.Walk(\"internal/mymvc/views/\"+dir, func(path string, info os.FileInfo, err error) error {\n\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\n\t\tif !info.IsDir() && (strings.HasSuffix(path, \".html\") || strings.HasSuffix(path, \".js\")) {\n\t\t\tfile, err := ioutil.ReadFile(path)\n\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\n\t\t\tfilename := strings.Replace(path, \"views/\"+dir, \"\", -1)\n\n\t\t\tif strings.Contains(filename, fileName) {\n\t\t\t\ttmpl = tmpl.New(filename)\n\n\t\t\t\ttmpl, err = tmpl.Parse(string(file) + layout)\n\n\t\t\t\tif err != nil {\n\t\t\t\t\treturn err\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\n\t\treturn nil\n\t}); err != nil {\n\t\treturn nil, err\n\t}\n\n\treturn tmpl, nil\n}\n-- internal/@@.ImportPath@@/controller/sample_controller.go --\npackage controller\n\nimport (\n\t\"encoding/json\"\n\t\"log\"\n\t\"net/http\"\n\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/model\"\n)\n\ntype SampleController struct {\n\tmodel.Model\n}\n\nfunc NewSampleController(sm model.Model) SampleController {\n\treturn SampleController{sm}\n}\n\nfunc (s SampleController) SampleIndex(w http.ResponseWriter, r *http.Request) {\n\tsamples, err := s.SampleModel.Fetch()\n\n\tif err != nil {\n\t\thttp.Error(w, err.Error(), http.StatusNotFound)\n\t}\n\n\tresp := &response{\n\t\tStatus: http.StatusOK,\n\t\tResult: samples,\n\t}\n\n\tres, err := json.Marshal(resp)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\t_, _ = w.Write(res)\n}\n\nfunc (s SampleController) SampleHTML(w http.ResponseWriter, r *http.Request) {\n\ttmpl, err := parseTemplate(\"sample\", \"index\")\n\n\tif err != nil {\n\t\tlog.Fatal(\"err :\", err)\n\t}\n\n\tif err := tmpl.Execute(w, nil); err != nil {\n\t\tlog.Printf(\"failed to execute template: %v\", err)\n\t}\n}\n-- internal/@@.ImportPath@@/model/model_base.go --\npackage model\n\nimport \"database/sql\"\n\ntype Model struct {\n\tSampleModel\n}\n\nfunc NewModel(conn *sql.DB) Model {\n\tm := Model{}\n\tm.SampleModel = NewSampleModel(conn)\n\treturn m\n}\n-- internal/@@.ImportPath@@/model/sample_model.go --\npackage model\n\nimport \"database/sql\"\n\ntype SampleModel struct {\n\tconn *sql.DB\n}\n\nfunc NewSampleModel(conn *sql.DB) SampleModel {\n\treturn SampleModel{conn}\n}\n\ntype Sample struct {\n\tID   int64 `json:\"id\" db:\"id\"`\n\tText string `json:\"text\" db:\"text\"`\n}\n\nfunc (s *SampleModel) Fetch() ([]*Sample, error) {\n\tvar samples []*Sample\n\trows, err := s.conn.Query(\"SELECT id, text FROM samples;\")\n\tif rows == nil {\n\t\treturn nil, err\n\t}\n\tfor rows.Next() {\n\t\tsample := &Sample{}\n\t\terr = rows.Scan(&sample.ID, &sample.Text)\n\t\tif err == nil {\n\t\t\tsamples = append(samples, sample)\n\t\t}\n\t}\n\treturn samples, err\n}\n-- internal/@@.ImportPath@@/router/router.go --\npackage router\n\nimport (\n\t\"github.com/go-chi/chi\"\n\t\"github.com/go-chi/chi/middleware\"\n\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/controller\"\n)\n\ntype Server struct {\n\tRoute *chi.Mux\n}\n\nfunc NewRouter() *Server {\n\treturn &Server{\n\t\tRoute: chi.NewRouter(),\n\t}\n}\n\nfunc (s *Server) Router(c controller.Controller) {\n\ts.Route.Use(middleware.Logger)\n\ts.Route.Use(middleware.Recoverer)\n\ts.Route.Route(\"/\", func(r chi.Router) {\n\t\tr.Get(\"/\", c.SampleController.SampleHTML)\n\t})\n\ts.Route.Route(\"/api/v1\", func(r chi.Router) {\n\t\tr.Get(\"/\", c.SampleController.SampleIndex)\n\t\t// TODO\n\t})\n}\n-- internal/@@.ImportPath@@/views/layout/_footer.html --\n{{ define \"footer\" }}\n    <footer>sample</footer>\n    </body>\n    </html>\n{{ end }}\n-- internal/@@.ImportPath@@/views/layout/_header.html --\n{{ define \"header\" }}\n    <!DOCTYPE html>\n    <html lang=\"en\">\n    <head>\n        <meta charset=\"UTF-8\">\n        <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n        <title>Sample</title>\n        <style>\n            * {\n                margin: 0;\n                padding: 0;\n                box-sizing: border-box;\n            }\n        </style>\n    </head>\n    <body>\n    <header><a href=\"/\">sample</a></header>\n{{ end }}\n-- internal/@@.ImportPath@@/views/sample/index.html --\n{{ template \"header\" }}\n<h1><i>SPRINTER</i></h1>\n{{ template \"footer\" }}\n"))
var MinimumTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage main\n\nimport (\n\t\"fmt\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/config/database/conf\"\n\t\"@@.ImportPath@@/internal/@@.ImportPath@@/router\"\n\t\"net/http\"\n\t\"os\"\n)\n\nconst defaultPort = \"8080\"\n\nfunc main() {\n\tport := os.Getenv(\"PORT\")\n\tif port == \"\" {\n\t\tport = defaultPort\n\t}\n\n\tconn, err := conf.NewDatabaseConnection()\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\tif conn == nil {\n\t\tpanic(err)\n\t}\n\tdefer func() {\n\t\tif conn != nil {\n\t\t\tif err := conn.Close(); err != nil {\n\t\t\t\tpanic(err)\n\t\t\t}\n\t\t}\n\t}()\n\n\tfmt.Println(`\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * \n  *        ####    #####    #####     ####    ##  ##   ######   ######   #####  *\n  *      ##  ##   ##  ##   ##  ##     ##     ### ##     ##     ##       ##  ##  *\n  *     ##       ##  ##   ##  ##     ##     ######     ##     ##       ##  ##   *\n  *     ####    #####    #####      ##     ######     ##     ####     #####     *\n  *       ##   ##       ####       ##     ## ###     ##     ##       ####       *\n  *  ##  ##   ##       ## ##      ##     ##  ##     ##     ##       ## ##       *\n  *  ####    ##       ##  ##    ####    ##  ##     ##     ######   ##  ##       *\n    * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n\n    API:\tGET http://localhost:8080/api/v1\n    `)\n\n\ts := router.NewRouter(conn)\n\ts.Router()\n\n\t_ = http.ListenAndServe(\":\"+port, s.Route)\n}\n-- internal/@@.ImportPath@@/config/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n-- internal/@@.ImportPath@@/router/router.go --\npackage router\n\nimport (\n\t\"database/sql\"\n\t\"encoding/json\"\n\t\"github.com/go-chi/chi\"\n\t\"github.com/go-chi/chi/middleware\"\n\t\"log\"\n\t\"net/http\"\n)\n\ntype Server struct {\n\tRoute *chi.Mux\n\tConn *sql.DB\n}\n\nfunc NewRouter(conn *sql.DB) *Server {\n\treturn &Server{\n\t\tRoute: chi.NewRouter(),\n\t\tConn: conn,\n\t}\n}\n\ntype Sample struct {\n\tID   int64 `json:\"id\" db:\"id\"`\n\tText string `json:\"text\" db:\"text\"`\n}\n\nfunc (s *Server) Router() {\n\ts.Route.Use(middleware.Logger)\n\ts.Route.Use(middleware.Recoverer)\n\ts.Route.Route(\"/api/v1\", func(r chi.Router) {\n\t\tr.Get(\"/\", func(w http.ResponseWriter, r *http.Request) {\n\t\t\tvar samples []*Sample\n\t\t\trows, err := s.Conn.Query(\"SELECT id, text FROM samples;\")\n\t\t\tif rows == nil {\n\t\t\t\tlog.Fatalln(err)\n\t\t\t}\n\t\t\tfor rows.Next() {\n\t\t\t\tsample := &Sample{}\n\t\t\t\terr = rows.Scan(&sample.ID, &sample.Text)\n\t\t\t\tif err == nil {\n\t\t\t\t\tsamples = append(samples, sample)\n\t\t\t\t}\n\t\t\t}\n\n\t\t\tres, err := json.Marshal(samples)\n\t\t\tif err != nil {\n\t\t\t\tlog.Fatal(err)\n\t\t\t}\n\n\t\t\t_, _ = w.Write(res)\n\t\t})\n\t})\n}\n"))
var CleanTmpl = template.Must(template.New("tmpl").Delims(`@@`, `@@`).Parse(layout+"-- cmd/main.go --\npackage cmd\n-- internal/@@.ImportPath@@/infrastructure/database/conf/database.go --\npackage conf\n\nimport (\n\t\"database/sql\"\n\t@@ if eq .DataBase \"MySQL\" -@@\n\t_ \"github.com/go-sql-driver/mysql\"\n\t@@ else @@\n\t_ \"github.com/lib/pq\"\n\t@@ end @@\n\t\"os\"\n)\n\nvar (\n\tsource = os.Getenv(\"DATABASE_URL\")\n\tdriver = os.Getenv(\"DRIVER\")\n)\n\nfunc NewDatabaseConnection() (*sql.DB, error) {\n\tconn, err := sql.Open(driver, source)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn conn, nil\n}\n"))
