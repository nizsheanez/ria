FROM nizsheanez/nodejs

RUN npm install -g gulp gulp-jade gulp-stylus gulp-livereload gulp-myth gulp-csso gulp-imagemin gulp-uglify gulp-concat connect --save-dev

RUN mkdir /gulp

WORKDIR /gopath/src/ria/static

RUN

CMD ["gulp", "watch"]
