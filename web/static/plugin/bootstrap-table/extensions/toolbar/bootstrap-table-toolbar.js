const Utils=$.fn.bootstrapTable.utils,bootstrap={bootstrap3:{icons:{advancedSearchIcon:"glyphicon-chevron-down"},html:{modal:`
        <div id="avdSearchModal_%s"  class="modal fade" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel" aria-hidden="true">
          <div class="modal-dialog modal-xs">
            <div class="modal-content">
              <div class="modal-header">
                <h4 class="modal-title">%s</h4>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                  <span aria-hidden="true">&times;</span>
                </button>
              </div>
              <div class="modal-body modal-body-custom">
                <div class="container-fluid" id="avdSearchModalContent_%s"
                  style="padding-right: 0px; padding-left: 0px;" >
                </div>
              </div>
              <div class="modal-footer">
                <button type="button" id="btnCloseAvd_%s" class="btn btn-%s">%s</button>
              </div>
            </div>
          </div>
        </div>
      `}},bootstrap4:{icons:{advancedSearchIcon:"fa-chevron-down"},html:{modal:`
        <div id="avdSearchModal_%s"  class="modal fade" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel" aria-hidden="true">
          <div class="modal-dialog modal-xs">
            <div class="modal-content">
              <div class="modal-header">
                <h4 class="modal-title">%s</h4>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                  <span aria-hidden="true">&times;</span>
                </button>
              </div>
              <div class="modal-body modal-body-custom">
                <div class="container-fluid" id="avdSearchModalContent_%s"
                  style="padding-right: 0px; padding-left: 0px;" >
                </div>
              </div>
              <div class="modal-footer">
                <button type="button" id="btnCloseAvd_%s" class="btn btn-%s">%s</button>
              </div>
            </div>
          </div>
        </div>
      `}},bulma:{icons:{advancedSearchIcon:"fa-chevron-down"},html:{modal:`
        <div class="modal" id="avdSearchModal_%s">
          <div class="modal-background"></div>
          <div class="modal-card">
            <header class="modal-card-head">
              <p class="modal-card-title">%s</p>
              <button class="delete" aria-label="close"></button>
            </header>
            <section class="modal-card-body" id="avdSearchModalContent_%s"></section>
            <footer class="modal-card-foot">
              <button class="button" id="btnCloseAvd_%s" data-close="btn btn-%s">%s</button>
            </footer>
          </div>
        </div>
      `}},foundation:{icons:{advancedSearchIcon:"fa-chevron-down"},html:{modal:`
        <div class="reveal" id="avdSearchModal_%s" data-reveal>
          <h1>%s</h1>
          <div id="avdSearchModalContent_%s">
          
          </div>
          <button class="close-button" data-close aria-label="Close modal" type="button">
            <span aria-hidden="true">&times;</span>
          </button>
          
          <button id="btnCloseAvd_%s" class="%s" type="button">%s</button>
        </div>
      `}},materialize:{icons:{advancedSearchIcon:"expand_more"},html:{modal:`
        <div id="avdSearchModal_%s" class="modal">
          <div class="modal-content">
            <h4>%s</h4>
            <div id="avdSearchModalContent_%s">
            
            </div>
          </div>
          <div class="modal-footer">
            <a href="javascript:void(0)"" id="btnCloseAvd_%s" class="modal-close waves-effect waves-green btn-flat %s">%s</a>
          </div>
        </div>
      `}},semantic:{icons:{advancedSearchIcon:"fa-chevron-down"},html:{modal:`
        <div class="ui modal" id="avdSearchModal_%s">
          <i class="close icon"></i>
          <div class="header">
            %s
          </div>
          <div class="image content ui form" id="avdSearchModalContent_%s"></div>
          <div class="actions">
            <div id="btnCloseAvd_%s" class="ui black deny button %s">%s</div>
          </div>
        </div>
      `}}}[$.fn.bootstrapTable.theme];$.extend($.fn.bootstrapTable.defaults,{advancedSearch:!1,idForm:"advancedSearch",actionForm:"",idTable:void 0,onColumnAdvancedSearch(a,t){return!1}}),$.extend($.fn.bootstrapTable.defaults.icons,{advancedSearchIcon:bootstrap.icons.advancedSearchIcon}),$.extend($.fn.bootstrapTable.Constructor.EVENTS,{"column-advanced-search.bs.table":"onColumnAdvancedSearch"}),$.extend($.fn.bootstrapTable.locales,{formatAdvancedSearch(){return"Advanced search"},formatAdvancedCloseButton(){return"Close"}}),$.extend($.fn.bootstrapTable.defaults,$.fn.bootstrapTable.locales),$.BootstrapTable=class extends $.BootstrapTable{initToolbar(){const a=this.options;this.showToolbar=this.showToolbar||a.search&&a.advancedSearch&&a.idTable,super.initToolbar(),!(!a.search||!a.advancedSearch||!a.idTable)&&(this.$toolbar.find(">.columns").append(`
      <button class="${this.constants.buttonsClass} "
        type="button"
        name="advancedSearch"
        aria-label="advanced search"
        title="${a.formatAdvancedSearch()}">
        ${this.options.showButtonIcons?Utils.sprintf(this.constants.html.icon,a.iconsPrefix,a.icons.advancedSearchIcon):""}
        ${this.options.showButtonText?this.options.formatAdvancedSearch():""}
      </button>
    `),this.$toolbar.find('button[name="advancedSearch"]').off("click").on("click",()=>this.showAvdSearch()))}showAvdSearch(){const a=this.options,t="#avdSearchModal_"+a.idTable;if($(t).length<=0){$("body").append(Utils.sprintf(bootstrap.html.modal,a.idTable,a.formatAdvancedSearch(),a.idTable,a.idTable,a.buttonsClass,a.formatAdvancedCloseButton()));let o=0;$(`#avdSearchModalContent_${a.idTable}`).append(this.createFormAvd().join("")),$(`#${a.idForm}`).off("keyup blur","input").on("keyup blur","input",s=>{a.sidePagination==="server"?this.onColumnAdvancedSearch(s):(clearTimeout(o),o=setTimeout(()=>{this.onColumnAdvancedSearch(s)},a.searchTimeOut))}),$(`#btnCloseAvd_${a.idTable}`).click(()=>this.hideModal()),$.fn.bootstrapTable.theme==="bulma"&&$(t).find(".delete").off("click").on("click",()=>this.hideModal()),this.showModal()}else this.showModal()}showModal(){const a="#avdSearchModal_"+this.options.idTable;$.inArray($.fn.bootstrapTable.theme,["bootstrap3","bootstrap4"])!==-1?$(a).modal():$.fn.bootstrapTable.theme==="bulma"?$(a).toggleClass("is-active"):$.fn.bootstrapTable.theme==="foundation"?(this.toolbarModal||(this.toolbarModal=new Foundation.Reveal($(a))),this.toolbarModal.open()):$.fn.bootstrapTable.theme==="materialize"?($(a).modal(),$(a).modal("open")):$.fn.bootstrapTable.theme==="semantic"&&$(a).modal("show")}hideModal(){const a=$(`#avdSearchModal_${this.options.idTable}`),t="#avdSearchModal_"+this.options.idTable;$.inArray($.fn.bootstrapTable.theme,["bootstrap3","bootstrap4"])!==-1?a.modal("hide"):$.fn.bootstrapTable.theme==="bulma"?($("html").toggleClass("is-clipped"),$(t).toggleClass("is-active")):$.fn.bootstrapTable.theme==="foundation"?this.toolbarModal.close():$.fn.bootstrapTable.theme==="materialize"?$(t).modal("open"):$.fn.bootstrapTable.theme==="semantic"&&$(t).modal("close"),this.options.sidePagination==="server"&&(this.options.pageNumber=1,this.updatePagination(),this.trigger("column-advanced-search",this.filterColumnsPartial))}createFormAvd(){const a=this.options,t=[`<form class="form-horizontal" id="${a.idForm}" action="${a.actionForm}">`];for(const o of this.columns)!o.checkbox&&o.visible&&o.searchable&&t.push(`
          <div class="form-group row">
            <label class="col-sm-4 control-label">${o.title}</label>
            <div class="col-sm-6">
              <input type="text" class="form-control ${this.constants.classes.input}" name="${o.field}" placeholder="${o.title}" id="${o.field}">
            </div>
          </div>
        `);return t.push("</form>"),t}initSearch(){if(super.initSearch(),!this.options.advancedSearch||this.options.sidePagination==="server")return;const a=$.isEmptyObject(this.filterColumnsPartial)?null:this.filterColumnsPartial;this.data=a?this.data.filter((t,o)=>{for(const[s,l]of Object.entries(a)){const i=l.toLowerCase();let e=t[s];const d=this.header.fields.indexOf(s);if(e=Utils.calculateObjectValue(this.header,this.header.formatters[d],[e,t,o],e),!(d!==-1&&(typeof e=="string"||typeof e=="number")&&`${e}`.toLowerCase().includes(i)))return!1}return!0}):this.data}onColumnAdvancedSearch(a){const t=$.trim($(a.currentTarget).val()),o=$(a.currentTarget)[0].id;$.isEmptyObject(this.filterColumnsPartial)&&(this.filterColumnsPartial={}),t?this.filterColumnsPartial[o]=t:delete this.filterColumnsPartial[o],this.options.sidePagination!=="server"&&(this.options.pageNumber=1,this.onSearch(a),this.updatePagination(),this.trigger("column-advanced-search",o,t))}};
