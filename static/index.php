<?php
require(__DIR__ . '/../../common/components/debug.php');

ini_set("display_errors", 1);
error_reporting(E_ALL);
ini_set('xdebug.max_nesting_level', 1000);

ini_set('session.cookie_lifetime', 60 * 60 * 24 * 30);

// comment out the following line to disable debug mode
defined('YII_DEBUG') or define('YII_DEBUG', true);
defined('YII_ENV') or define('YII_ENV', 'dev');

require(__DIR__ . '/../../vendor/autoload.php');
require(__DIR__ . '/../../vendor/yiisoft/yii2/Yii.php');
require(__DIR__ . '/../../common/config/aliases.php');

$config = require(__DIR__ . '/../config/main.php');
(new yii\web\Application($config))->run();
