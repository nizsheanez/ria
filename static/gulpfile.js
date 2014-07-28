// include gulp
var gulp = require('gulp'),
    concat = require('gulp-concat'),
    jshint = require('gulp-jshint'),
    less = require('gulp-less'),
    rename = require('gulp-rename'),
    rimraf = require('gulp-rimraf'),
    minifycss = require('gulp-minify-css'),
    imagemin = require('gulp-imagemin'),
    path = require('path'),
    ngannotate = require('gulp-ng-annotate'),
    watch = require('gulp-watch'),
    uglify = require('gulp-uglify'),
    shell = require('gulp-shell'),
    livereload = require('gulp-livereload'),
    debug = require('gulp-debug');

var conf = {
    app: './src/',
    dist: './assets/build/',
    root: '../../'
};

var js = [
        conf.app + '/vendor/jquery/dist/jquery.min.js',
//    conf.app + '/vendor/jquery-ui/ui/minified/jquery-ui.min.js',
        conf.dist + '/ngmin/vendor/angular.min.js',
        conf.app + '/vendor/angular-resource/angular-resource.min.js',
        conf.app + '/vendor/angular-route/angular-route.min.js',
        conf.app + '/vendor/angular-translate/*.min.js',
        conf.app + '/vendor/angular-ui/*.min.js',
        conf.app + '/vendor/angular-bootstrap/ui-bootstrap.min.js',
        conf.app + '/vendor/angular-bootstrap/ui-bootstrap-tpls.js',
        conf.app + '/vendor/angular-sortable/src/sortable.js',
        conf.dist + '/ngmin/vendor/angular-elastic/elastic.js',
        conf.app + '/vendor/angular-ui-utils/*.min.js',
        conf.app + '/vendor/angular-sanitize/*.min.js',
        conf.app + '/vendor/textAngular/*.min.js',


        conf.dist + '/ngmin/websocket/**/*.js',

        conf.dist + '/ngmin/common/**/*.js',
        conf.dist + '/ngmin/app/app.js',
        conf.dist + '/ngmin/app/goal/services/goal.js',
        conf.dist + '/ngmin/app/goal/services/modal.js',
        conf.dist + '/ngmin/app/goal/services/category.js',
        conf.dist + '/ngmin/app/goal/services/report.js',
        conf.dist + '/ngmin/app/goal/services/tpl.js',
        conf.dist + '/ngmin/app/goal/services/user.js',
        conf.dist + '/ngmin/app/goal/services/server.js',
        conf.dist + '/ngmin/app/goal/services/alert.js',

        conf.dist + '/ngmin/app/goal/controllers/*.js',
        conf.dist + '/ngmin/app/goal/directives/*.js'
];

gulp.task('fonts', function () {
    gulp.src(conf.app + '/vendor/components-font-awesome/fonts/*')
        .pipe(gulp.dest(conf.dist + '/fonts/'));
});

gulp.task('js.copy', function () {

    gulp.src(conf.app + '/vendor/angular-route/angular-route*')
        .pipe(gulp.dest(conf.dist));

    gulp.src(conf.app + '/vendor/jquery/dist/jquery.min.map')
        .pipe(gulp.dest(conf.dist));

    gulp.src(conf.app + '/vendor/angular/*.map')
        .pipe(gulp.dest(conf.dist));


});

gulp.task('js.ngmin', function () {

    gulp.src(conf.app + '/app/**/*.js')
        .pipe(ngannotate())
        .pipe(gulp.dest(conf.dist + '/ngmin/app'));

    gulp.src([conf.app + '/vendor/angular-elastic/elastic.js', conf.app + '/vendor/angular/angular.min.js'])
        .pipe(ngannotate())
        .pipe(gulp.dest(conf.dist + '/ngmin/vendor'));

    gulp.src([conf.app + '/common/**/*.js'])
        .pipe(ngannotate())
        .pipe(gulp.dest(conf.dist + '/ngmin/common'));

    gulp.src(conf.app + '/../../../vendor/nizsheanez/yii2-websocket-application/src/wamp/assets/js/*')
        .pipe(ngannotate())
        .pipe(gulp.dest(conf.dist + "ngmin/websocket"));


});

gulp.task('js.concat', ['js.ngmin', 'js.copy'], function () {

    gulp.src(js)
        .pipe(debug({verbose: false}))
        .pipe(concat('all.js'))
        .pipe(gulp.dest(conf.dist))
});

gulp.task('js.compress', ['js.concat'], function () {

    return gulp.src(conf.dist + '/all.js')
        .pipe(debug({verbose: false}))
        .pipe(uglify({outSourceMap: true}))
        .pipe(gulp.dest(conf.dist + '/min'));
});


gulp.task('less', function () {
    return gulp.src(conf.app + '/less/site.less')
        .pipe(less({
            paths: [ path.join(__dirname, 'less', 'includes') ]
        }))
        .pipe(concat('site.css'))
        .pipe(gulp.dest(conf.dist + '/css'))
        ;
});

gulp.task('clean', function () {
    return gulp.src([conf.dist], {read: false})
        .pipe(rimraf());
});

gulp.task('js', ['js.compress']);
gulp.task('js.dev', ['js.concat'], function () {
    gulp.src('');
});

gulp.task('build', ['clean'], function () {
    gulp.start('js', 'less', 'fonts');
});
gulp.task('build.dev', ['clean'], function () {
    gulp.start('fonts', 'js.dev', 'less');
});

gulp.task('watch', ['build.dev'], function () {

    livereload.listen();

    gulp.watch(conf.app + '/**/*.js', ['js.dev']).on('change', livereload.changed);
    gulp.watch(conf.app + 'less/**/*', ['less']).on('change', livereload.changed);

//        gulp.watch([conf.app + '/**/*', conf.root + '/views/layouts/main.php', conf.app + '../less/**/*'], ['livereload']);

});