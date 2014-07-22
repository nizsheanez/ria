'use strict';

module.exports = function (grunt) {

    // Load grunt tasks automatically
    require('load-grunt-tasks')(grunt);

    // Time how long tasks take. Can help when optimizing build times
    require('time-grunt')(grunt);

    // Define the configuration for all the tasks
    grunt.initConfig({

        // Project settings
        conf: {
            // configurable paths
            app: 'src/',
            dist: 'assets/build/',
            root: '../../'
        },

        // Watches files for changes and runs tasks based on the changed files
        watch: {
            js: {
                files: ['<%= conf.app %>/**/*.js'],
                tasks: [
                    'ngmin', 'concat'
                ],
                options: {
                    livereload: true
                }
            },
//            jsTest: {
//                files: ['test/spec/{,*/}*.js'],
//                tasks: ['newer:jshint:test', 'karma']
//            },
            phpTests: {
                files: ['<%= conf.root %>frontend/**/*.php', '<%= conf.root %>common/**/*.php'],
                tasks: [
                    'shell:phpUnitTests'
                ]
            },
            css: {
                files: ['<%= conf.app %>/less/**/*'],
                tasks: ['less'],
                options: {
                    livereload: true
                }
            },
            html: {
                    files: ['<%= conf.app %>/**/*.html'],
                tasks: [],
                options: {
                    livereload: true
                }
            },
            gruntfile: {
                files: ['Gruntfile.js']
            },
            livereload: {
                options: {
                    livereload: '<%= connect.options.livereload %>'
                },
                files: [
                    '<%= conf.app %>/**/*',
                    '<%= conf.app %>/../../views/layouts/main.php',
                    '<%= conf.app %>/../less/**/*'
                ]
            }
        },

        shell: {                                // Task
            phpUnitTests: {                      // Target
                options: {                      // Options
                    stdout: true,
                    stderr: true,
                    execOptions: {
                        cwd: '<%= conf.root %>frontend'
                    }
                },
                command: 'php ../vendor/bin/codecept run unit | grep -v " Ok"'
            }
        },

        // The actual grunt server settings
        connect: {
            options: {
                port: 9000,
                // Change this to '0.0.0.0' to access the server from outside.
                hostname: 'localhost',
                livereload: 35729
            },
            livereload: {
                options: {
                    open: true,
                    base: [
                        '.tmp',
                        '<%= conf.app %>'
                    ]
                }
            },
            test: {
                options: {
                    port: 9001,
                    base: [
                        '.tmp',
                        'test',
                        '<%= conf.app %>'
                    ]
                }
            },
            dist: {
                options: {
                    base: '<%= conf.dist %>'
                }
            }
        },

        // Empties folders to start fresh
        clean: {
            dist: {
                files: [
                    {
                        dot: true,
                        src: [
                            '.tmp',
                            '<%= conf.dist %>/*',
                            '!<%= conf.dist %>/.git*'
                        ]
                    }
                ]
            },
            server: '.tmp'
        },

        // Add vendor prefixed styles
        autoprefixer: {
            options: {
                browsers: ['last 1 version']
            },
            dist: {
                files: [
                    {
                        expand: true,
                        cwd: '.tmp/styles/',
                        src: '**/*.css',
                        dest: '.tmp/styles/'
                    }
                ]
            }
        },


        // Renames files for browser caching purposes
        rev: {
            dist: {
                files: {
                    src: [
                        '<%= conf.dist %>/scripts/**/*.js',
                        '<%= conf.dist %>/styles/**/*.css',
                        '<%= conf.dist %>/images/**/*.{png,jpg,jpeg,gif,webp,svg}',
                        '<%= conf.dist %>/styles/fonts/*'
                    ]
                }
            }
        },

        // Reads HTML for usemin blocks to enable smart builds that automatically
        // concat, minify and revision files. Creates configurations in memory so
        // additional tasks can operate on them
        useminPrepare: {
            html: '<%= conf.app %>/index.html',
            options: {
                dest: '<%= conf.dist %>'
            }
        },

        // Performs rewrites based on rev and the useminPrepare configuration
        usemin: {
            html: ['<%= conf.dist %>/**/*.html'],
            css: ['<%= conf.dist %>/styles/**/*.css'],
            options: {
                assetsDirs: ['<%= conf.dist %>']
            }
        },

        // The following *-min tasks produce minified files in the dist folder
        imagemin: {
            dist: {
                files: [
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/images',
                        src: '**/*.{png,jpg,jpeg,gif}',
                        dest: '<%= conf.dist %>/images'
                    }
                ]
            }
        },
        svgmin: {
            dist: {
                files: [
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/images',
                        src: '**/*.svg',
                        dest: '<%= conf.dist %>/images'
                    }
                ]
            }
        },
        htmlmin: {
            dist: {
                options: {
                    collapseWhitespace: true,
                    collapseBooleanAttributes: true,
                    removeCommentsFromCDATA: true,
                    removeOptionalTags: true
                },
                files: [
                    {
                        expand: true,
                        cwd: '<%= conf.dist %>',
                        src: ['*.html', 'views/**/*.html'],
                        dest: '<%= conf.dist %>'
                    }
                ]
            }
        },

        // Allow the use of non-minsafe AngularJS files. Automatically makes it
        // minsafe compatible so Uglify does not destroy the ng references
        ngmin: {
            dist: {
                files: [
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/app',
                        src: '**/*.js',
                        dest: '<%= conf.dist %>/concat/scripts/app'
                    },
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/vendor',
                        src: ['angular-elastic/elastic.js', 'angular/angular.min.js'],
                        dest: '<%= conf.dist %>/concat/scripts/vendor'
                    },
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/common/',
                        src: '**/*.js',
                        dest: '<%= conf.dist %>/concat/scripts/common'
                    }
                ]
            }
        },

        concat: {
            options: {
                mangle: {
                    separator: ';'
                }
            },
            dist: {
                files: {
                    '<%= conf.dist %>/all.js': [
                        '<%= conf.app %>/vendor/jquery/dist/jquery.min.js',
                        '<%= conf.app %>/vendor/jquery-ui/minified/jquery-ui.min.js',
                        '<%= conf.dist %>concat/scripts/vendor/angular/angular.min.js',
                        '<%= conf.app %>/vendor/angular-bootstrap/ui-bootstrap.min.js',
                        '<%= conf.app %>/vendor/angular-bootstrap/ui-bootstrap-tpls.js',
                        '<%= conf.app %>/vendor/angular-route/angular-route.min.js',
                        '<%= conf.app %>/vendor/angular-translate/*.min.js',
                        '<%= conf.app %>/vendor/angular-ui/*.min.js',
                        '<%= conf.app %>/vendor/angular-sortable/src/sortable.js',
                        '<%= conf.app %>/vendor/angular-ui-utils/*.min.js',
                        '<%= conf.app %>/vendor/angular-sanitize/*.min.js',
                        '<%= conf.app %>/vendor/textAngular/*.min.js',
                        '<%= conf.dist %>/concat/scripts/vendor/angular-elastic/elastic.js',


                        '<%= conf.dist %>/concat/scripts/common/**/*.js',
                        '<%= conf.dist %>/concat/scripts/app/app.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/goal.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/modal.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/category.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/report.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/tpl.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/user.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/server.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/services/alert.js',



//                        '<%= conf.dist %>/concat/scripts/app/goal/services/*.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/controllers/*.js',
                        '<%= conf.dist %>/concat/scripts/app/goal/directives/*.js'
                    ]
                }
            }
        },
        uglify: {
            options: {
                mangle: {
                    except: ['jQuery', 'Angular']
                }
            },
            dist: {
                files: {
                    '<%= conf.dist %>/all.min.js': ['<%= conf.dist %>/all.js']
                }
            }
        },

        // Copies remaining files to places other tasks can use
        copy: {
            dist: {
                files: [
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/vendor/components-font-awesome/fonts/',
                        src: '*',
                        dest: '<%= conf.dist %>/fonts/',
                        filter: 'isFile'
                    },
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/vendor/angular-route',
                        src: 'angular-route.*',
                        dest: '<%= conf.dist %>/',
                        filter: 'isFile'
                    },
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/vendor/jquery',
                        src: 'jquery.min.map',
                        dest: '<%= conf.dist %>/',
                        filter: 'isFile'
                    },
                    {
                        expand: true,
                        cwd: '<%= conf.app %>/vendor/angular',
                        src: '*.map',
                        dest: '<%= conf.dist %>/',
                        filter: 'isFile'
                    }
                ]
            }
        },

        // Run some tasks in parallel to speed up the build process
        concurrent: {
            server: [
            ],
            test: [
            ],
            dist: [
                'imagemin',
                'svgmin'
            ]
        },

        // By default, your `index.html`'s <!-- Usemin block --> will take care of
        // minification. These next options are pre-configured if you do not wish
        // to use the Usemin blocks.
        less: {
            development: {
                options: {
                    cleancss: true
                },
                files: {
                    "<%= conf.dist %>/css/site.css": "<%= conf.app %>/less/site.less"
                }
            },
            production: {
                options: {
                    cleancss: true
                },
                files: {
                    "<%= conf.dist %>/css/site.css": "<%= conf.app %>less/site.less"
                }
            }
        },

        // Test settings
        karma: {
            unit: {
                configFile: 'karma.conf.js',
                singleRun: true
            }
        }
    });


    grunt.registerTask('serve', function (target) {
        if (target === 'dist') {
            return grunt.task.run(['build', 'connect:dist:keepalive']);
        }

        grunt.task.run([
            'clean:server',
            'concurrent:server',
            'less',
            'copy',
            'ngmin',
            'concat',
//            'autoprefixer',
            'connect:livereload',
            'watch'
        ]);
    });

    grunt.registerTask('server', function () {
        grunt.log.warn('The `server` task has been deprecated. Use `grunt serve` to start a server.');
        grunt.task.run(['serve']);
    });

    grunt.registerTask('test', [
        'clean:server',
        'concurrent:test',
//        'autoprefixer',
        'connect:test'
//        'karma'
    ]);

    grunt.registerTask('build', [
        'clean:dist',
//        'useminPrepare',
//        'concurrent:dist',
//        'autoprefixer',
        'ngmin',
        'concat',
        'copy',
//        'copy:dist',
//        'copy:dist',
        'less',
        'uglify'
//        'rev',
//        'usemin',
//        'htmlmin'
    ]);

    grunt.registerTask('default', [
//    'newer:jshint',
        'test',
        'build'
    ]);
};
