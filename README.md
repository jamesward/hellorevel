Run Locally:

    go get github.com/robfig/revel
    go get github.com/jamesward/revel
    go build -o bin/revel github.com/jamesward/revel/cmd
    go get github.com/jamesward/hellorevel
    bin/revel run github.com/jamesward/hellorevel

Run on Heroku:

    heroku create --buildpack https://github.com/jamesward/heroku-buildpack-go-revel.git
    git push heroku master
    heroku open
