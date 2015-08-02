module.exports = function (grunt) {
  /*
   *  kanban/
   *    public/style.css <--
   *    assets/
   *      src/
   *        js/*.js
   *        scss/*.sass
   *      
   *
   *
   *
   */
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    /**
     * Sass
     */
    sass: {
      dev: {
        options: {
          style: 'expanded',
        },
        files: {
          'public/style.css': 'assets/src/scss/*.sass'
        }
      },
      heroku: {
        options: {
          style: 'compressed',
        },
        files: {
          '/app/public/style.css': 'assets/src/scss/*.sass'
        },
      }
    },
    /**
     * Watch
     */
    watch: {
      sass: {
        files: 'assets/src/scss/*.{scss,sass}',
        tasks: ['sass:dev']
      }
    }
  });
  require('matchdep').filterDev('grunt-*').forEach(grunt.loadNpmTasks);
  /**
   * Default task
   * Run `grunt` on the command line
   */
  grunt.registerTask('default', [
    'sass:dev',
    'watch'
  ]);
  grunt.registerTask('heroku', [
    'sass:heroku',
  ]);
};

