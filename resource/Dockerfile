FROM python:3.12-alpine
RUN apk add --no-cache build-base
WORKDIR /resource
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
CMD ["python", "src/main.py"]
