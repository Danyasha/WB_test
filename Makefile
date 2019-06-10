all:
	python3 install_db.py
	mv tv_storage.db DB/
	docker build -t go_app .