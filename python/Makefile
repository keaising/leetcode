.PHONY: fmt
all: fmt

fmt:
	black . -l 80 --exclude=venv

lint:
	pycodestyle . --exclude=venv  --max-line-length=80
