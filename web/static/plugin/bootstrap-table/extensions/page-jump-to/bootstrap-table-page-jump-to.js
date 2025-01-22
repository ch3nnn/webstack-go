const Utils=$.fn.bootstrapTable.utils;$.extend($.fn.bootstrapTable.defaults,{showJumpTo:!1}),$.extend($.fn.bootstrapTable.locales,{formatJumpTo(){return"GO"}}),$.extend($.fn.bootstrapTable.defaults,$.fn.bootstrapTable.locales),$.BootstrapTable=class extends $.BootstrapTable{initPagination(...o){if(super.initPagination(...o),this.options.showJumpTo){const s=this.$pagination.find("> .pagination");let t=s.find(".page-jump-to");t.length||(t=$(`
          <div class="page-jump-to ${this.constants.classes.inputGroup}">
          <input type="number" class="${this.constants.classes.input}${Utils.sprintf(" input-%s",this.options.iconSize)}" value="${this.options.pageNumber}">
          <button class="${this.constants.buttonsClass}"  type="button">
          ${this.options.formatJumpTo()}
          </button>
          </div>
        `).appendTo(s),t.on("click","button",n=>{this.selectPage(+$(n.target).parent(".page-jump-to").find("input").val())}))}}};
