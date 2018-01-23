<div class="container">
    <div class="row">
        <div class="col-xs-6 col-md-4 col-center-block">
            <h3>iOS应用OTA安装说明</h3>
            <span>首先需要安装SSL 证书到手机上
                <a title="iPhone" href="/download/myCA.cer">ssl 证书安装</a>
            </span>
            <button type="button"
                    class="btn btn-info"
                    onclick="location.href='itms-services://?action=download-manifest&url=https://test.me/manifest'">
                点击安装
            </button>
        </div>
    </div>
</div>