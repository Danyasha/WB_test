all:
	python3 install_db.py
	docker build -t go_app .