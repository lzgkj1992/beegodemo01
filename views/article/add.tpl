<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>增加文章</title>
</head>
<body>
    <form action="/article/doEdit" method="post">
        标题：<input type="text" name="title" />
        <br><br>
        内容：<input type="text" name="content" />
        <br>
        <br>
        <input type="submit" value="提交">
    </form>
</body>
</html>