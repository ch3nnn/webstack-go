package repository

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

const (
	DefaultAboutSite = "> ❤️ 基于 Golang 开源的网址导航网站项目，具备完整的前后台，您可以拿来制作自己平日收藏的网址导航。\n\n\n> 如果你也是开发者，如果你也正好喜欢折腾，那希望这个网站能给你带来一些作用。"
	DefaultAuthor    = `
<div class="col-sm-4">
    <div class="xe-widget xe-conversations box2 label-info" onclick="window.open('https://blog.ch3nnn.cn/about/', '_blank')" data-toggle="tooltip" data-placement="bottom" title="" data-original-title="https://blog.ch3nnn.cn/about/">
        <div class="xe-comment-entry">
            <a class="xe-user-img">
                <img src="https://s2.loli.net/2023/02/20/H1k52mlXNYKWDrU.png" class="img-circle" width="40">
            </a>
            <div class="xe-comment">
                <a href="#" class="xe-user-name overflowClip_1">
                    <strong>Developer. Ch3nnn.</strong>
                </a>
                <p class="overflowClip_2"> 折腾不息 · 乐此不疲.</p>
            </div>
        </div>
    </div>
</div>

<div class="col-md-8">
    <div class="row">
        <div class="col-sm-12">
            <br>
            <blockquote>
                <p>
                    这是一个公益项目，而且是<a href="https://github.com/ch3nnn/webstack-go"> 开源 </a>的。你也可以拿来制作自己的网址导航。如果你有更好的想法，可以通过个人网站<a href="https://ch3nnn.cn/about/">ch3nnn.cn</a>中的联系方式找到我，欢迎与我交流分享。
                </p>
            </blockquote>
        </div>
    </div>
    <br>
</div>
`
	DefaultSiteTitle   = "WebStack-Go - 网址导航"
	DefaultSiteKeyword = "网址导航"
	DefaultSiteDesc    = "WebStack-Go - 基于 Golang 开源的网址导航网站"
)

var _ iCustomGenSysConfigFunc = (*customSysConfigDao)(nil)

type (
	// ISysConfigDao not edit interface name
	ISysConfigDao interface {
		iWhereSysConfigFunc
		WithContext(ctx context.Context) iCustomGenSysConfigFunc

		// TODO Custom WhereFunc ....
		// ...
	}

	// not edit interface name
	iCustomGenSysConfigFunc interface {
		iGenSysConfigFunc

		// TODO Custom DaoFunc ....
		// ...
	}

	// not edit interface name
	customSysConfigDao struct {
		sysConfigDao
	}
)

func NewSysConfigDao() ISysConfigDao {
	return &customSysConfigDao{
		sysConfigDao{
			sysConfigDo: query.SysConfig.WithContext(context.Background()),
		},
	}
}

func (d *customSysConfigDao) WithContext(ctx context.Context) iCustomGenSysConfigFunc {
	d.sysConfigDo = d.sysConfigDo.WithContext(ctx)
	return d
}
