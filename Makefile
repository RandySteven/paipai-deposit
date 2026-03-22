include ./files/env/.env
export


yaml_file = ./files/yaml/app.local.yml
cmd_folder = ./cmd/
gorun = @go run

ifeq ($(ENV),prod)
	yaml_file = ./files/yaml/app.prod.yml
else ifeq ($(ENV),staging)
	yaml_file = ./files/yaml/app.docker.yml
else ifeq ($(ENV),dev)
	yaml_file = ./files/yaml/app.local.yml
else
	$(error unknown variable in .env file)
endif

run:
	${gorun} ${cmd_folder}main -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}

seed:
	${gorun} ${cmd_folder}seed -config ${yaml_file}

drop:
	${gorun} ${cmd_folder}drop -config ${yaml_file}

test_env:
	${yaml_file}

env_check:
	$(ENV)

refresh: drop migration seed

run-docker:
	docker compose up --build -d

stop-docker:
	docker compose down

define create_model
    $(eval MODELNAME := $(shell bash -c 'read -p "Model name : " modelfile; echo $$modelfile'))
    $(eval LOWER_FIRST_CHAR := $(shell echo $(MODELNAME) | cut -c1 | tr '[:upper:]' '[:lower:]'))
    $(eval UPPER_FIRST_CHAR := $(shell echo $(MODELNAME) | cut -c1 ))
    $(eval MODELFILE := $(subst $(UPPER_FIRST_CHAR),$(LOWER_FIRST_CHAR),$(MODELNAME)))
    @echo "Creating model file: $(MODELFILE).go"
    @echo "Creating repository file: $(MODELFILE)_repository.go"
    @echo "package model" > $(MODELFILE).go
    @echo "type $(MODELNAME) struct {}" >> $(MODELFILE).go
    @echo "package repositories" > $(MODELFILE)_repository.go
    @echo "" >> $(MODELFILE)_repository.go
    @echo "import \"models\"" >> $(MODELFILE)_repository.go
    @echo "" >> $(MODELFILE)_repository.go
    @echo "type I$(MODELNAME)Repository interface{" >> $(MODELFILE)_repository.go
    @echo "    IRepository[models.$(MODELNAME)]" >> $(MODELFILE)_repository.go
    @echo "}" >> $(MODELFILE)_repository.go
endef


.PHONY: make_model

make_model:
	$(call create_model)

gen_proto:
	protoc --go_out=./proto --go-grpc_out=./proto --proto_path=./proto ./proto/service.proto