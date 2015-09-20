{{template "base/base.tpl" .}}
{{define "main"}}
<div class="container">
    <div class="page-header text-center">
        <h1>注册</h1>
    </div>
    {{if .ValidateMessage}}
    <div class="alert alert-dismissible alert-info">
        <button type="button" class="close" data-dismiss="alert">×</button>
        {{ .ValidateMessage }}
    </div>
    {{end}}
    <form class="form-horizontal" method="POST" action="/user/signup">
        <div class="form-group">
            <label for="email" class="col-sm-2 control-label">Email</label>
            <div class="col-sm-10">
                <input type="email" class="form-control" id="email" name="email" placeholder="请输入 Email">
            </div>
        </div>
        <div class="form-group">
            <label for="username" class="col-sm-2 control-label">Username</label>
            <div class="col-sm-10">
                <input type="text" class="form-control" id="username" name="username" placeholder="请输入用户名">
            </div>
        </div>
        <div class="form-group">
            <label for="password" class="col-sm-2 control-label">Password</label>
            <div class="col-sm-10">
                <input type="password" class="form-control" id="password" name="password" placeholder="请输入密码">
            </div>
        </div>
        <div class="form-group">
            <label for="password2" class="col-sm-2 control-label">Password</label>
            <div class="col-sm-10">
                <input type="password2" class="form-control" id="password2" name="password2" placeholder="请再次输入密码">
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-offset-2 col-sm-10">
                <button type="submit" class="btn btn-default">注册</button>
            </div>
        </div>
    </form>
</div>
{{end}}
