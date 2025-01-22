let isSingleSort=!1;const Utils=$.fn.bootstrapTable.utils,bootstrap={bootstrap3:{icons:{plus:"glyphicon-plus",minus:"glyphicon-minus",sort:"glyphicon-sort"},html:{multipleSortModal:`
        <div class="modal fade" id="%s" tabindex="-1" role="dialog" aria-labelledby="%sLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                     <h4 class="modal-title" id="%sLabel">%s</h4>
                </div>
                <div class="modal-body">
                    <div class="bootstrap-table">
                        <div class="fixed-table-toolbar">
                            <div class="bars">
                                <div id="toolbar">
                                     <button id="add" type="button" class="btn btn-default">%s %s</button>
                                     <button id="delete" type="button" class="btn btn-default" disabled>%s %s</button>
                                </div>
                            </div>
                        </div>
                        <div class="fixed-table-container">
                            <table id="multi-sort" class="table">
                                <thead>
                                    <tr>
                                        <th></th>
                                         <th><div class="th-inner">%s</div></th>
                                         <th><div class="th-inner">%s</div></th>
                                    </tr>
                                </thead>
                                <tbody></tbody>
                            </table>
                        </div>
                    </div>
                </div>
                <div class="modal-footer">
                     <button type="button" class="btn btn-default" data-dismiss="modal">%s</button>
                     <button type="button" class="btn btn-primary multi-sort-order-button">%s</button>
                </div>
            </div>
        </div>
    </div>
      `,multipleSortButton:'<button class="multi-sort btn btn-default" type="button" data-toggle="modal" data-target="#%s" title="%s">%s</button>',multipleSortSelect:'<select class="%s %s form-control">'}},bootstrap4:{icons:{plus:"fa-plus",minus:"fa-minus",sort:"fa-sort"},html:{multipleSortModal:`
        <div class="modal fade" id="%s" tabindex="-1" role="dialog" aria-labelledby="%sLabel" aria-hidden="true">
          <div class="modal-dialog" role="document">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title" id="%sLabel">%s</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                  <span aria-hidden="true">&times;</span>
                </button>
              </div>
              <div class="modal-body">
                <div class="bootstrap-table">
                        <div class="fixed-table-toolbar">
                            <div class="bars">
                                <div id="toolbar" class="pb-3">
                                     <button id="add" type="button" class="btn btn-secondary">%s %s</button>
                                     <button id="delete" type="button" class="btn btn-secondary" disabled>%s %s</button>
                                </div>
                            </div>
                        </div>
                        <div class="fixed-table-container">
                            <table id="multi-sort" class="table">
                                <thead>
                                    <tr>
                                        <th></th>
                                         <th><div class="th-inner">%s</div></th>
                                         <th><div class="th-inner">%s</div></th>
                                    </tr>
                                </thead>
                                <tbody></tbody>
                            </table>
                        </div>
                    </div>
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">%s</button>
                <button type="button" class="btn btn-primary multi-sort-order-button">%s</button>
              </div>
            </div>
          </div>
        </div>
      `,multipleSortButton:'<button class="multi-sort btn btn-secondary" type="button" data-toggle="modal" data-target="#%s" title="%s">%s</button>',multipleSortSelect:'<select class="%s %s form-control">'}},semantic:{icons:{plus:"fa-plus",minus:"fa-minus",sort:"fa-sort"},html:{multipleSortModal:`
        <div class="ui modal tiny" id="%s" aria-labelledby="%sLabel" aria-hidden="true">
        <i class="close icon"></i>
        <div class="header" id="%sLabel">
          %s
        </div>
        <div class="image content">
          <div class="bootstrap-table">
            <div class="fixed-table-toolbar">
                <div class="bars">
                  <div id="toolbar" class="pb-3">
                    <button id="add" type="button" class="ui button">%s %s</button>
                    <button id="delete" type="button" class="ui button" disabled>%s %s</button>
                  </div>
                </div>
            </div>
            <div class="fixed-table-container">
                <table id="multi-sort" class="table">
                    <thead>
                        <tr>
                            <th></th>
                            <th><div class="th-inner">%s</div></th>
                            <th><div class="th-inner">%s</div></th>
                        </tr>
                    </thead>
                    <tbody></tbody>
                </table>
            </div>
          </div>
        </div>
        <div class="actions">
          <div class="ui button deny">%s</div>
          <div class="ui button approve multi-sort-order-button">%s</div>
        </div>
      </div>
      `,multipleSortButton:'<button class="multi-sort ui button" type="button" data-toggle="modal" data-target="#%s" title="%s">%s</button>',multipleSortSelect:'<select class="%s %s">'}},materialize:{icons:{plus:"plus",minus:"minus",sort:"sort"},html:{multipleSortModal:`
        <div id="%s" class="modal" aria-labelledby="%sLabel" aria-hidden="true">
          <div class="modal-content" id="%sLabel">
            <h4>%s</h4>
            <div class="bootstrap-table">
            <div class="fixed-table-toolbar">
                <div class="bars">
                  <div id="toolbar" class="pb-3">
                    <button id="add" type="button" class="waves-effect waves-light btn">%s %s</button>
                    <button id="delete" type="button" class="waves-effect waves-light btn" disabled>%s %s</button>
                  </div>
                </div>
            </div>
            <div class="fixed-table-container">
                <table id="multi-sort" class="table">
                    <thead>
                        <tr>
                            <th></th>
                            <th><div class="th-inner">%s</div></th>
                            <th><div class="th-inner">%s</div></th>
                        </tr>
                    </thead>
                    <tbody></tbody>
                </table>
            </div>
          </div>
          <div class="modal-footer">
            <a href="javascript:void(0)" class="modal-close waves-effect waves-light btn">%s</a>
            <a href="javascript:void(0)" class="modal-close waves-effect waves-light btn multi-sort-order-button">%s</a>
          </div>
          </div>
        </div>
      `,multipleSortButton:'<a href="#%s" class="multi-sort waves-effect waves-light btn modal-trigger" type="button" data-toggle="modal" title="%s">%s</a>',multipleSortSelect:'<select class="%s %s browser-default">'}},foundation:{icons:{plus:"fa-plus",minus:"fa-minus",sort:"fa-sort"},html:{multipleSortModal:`
        <div class="reveal" id="%s" data-reveal aria-labelledby="%sLabel" aria-hidden="true">
            <div id="%sLabel">
              <h1>%s</h1>
              <div class="bootstrap-table">
                <div class="fixed-table-toolbar">
                    <div class="bars">
                      <div id="toolbar" class="padding-bottom-2">
                        <button id="add" type="button" class="waves-effect waves-light button">%s %s</button>
                        <button id="delete" type="button" class="waves-effect waves-light button" disabled>%s %s</button>
                      </div>
                    </div>
                </div>
                <div class="fixed-table-container">
                    <table id="multi-sort" class="table">
                        <thead>
                            <tr>
                                <th></th>
                                <th><div class="th-inner">%s</div></th>
                                <th><div class="th-inner">%s</div></th>
                            </tr>
                        </thead>
                        <tbody></tbody>
                    </table>
                </div>
              </div>

              <button class="waves-effect waves-light button" data-close aria-label="Close modal" type="button">
                <span aria-hidden="true">%s</span>
              </button>
              <button class="waves-effect waves-light button multi-sort-order-button" data-close aria-label="Order" type="button">
                  <span aria-hidden="true">%s</span>
              </button>
            </div>
        </div>
      `,multipleSortButton:'<button class="button multi-sort" data-open="%s" title="%s">%s</button>',multipleSortSelect:'<select class="%s %s browser-default">'}},bulma:{icons:{plus:"fa-plus",minus:"fa-minus",sort:"fa-sort"},html:{multipleSortModal:`
        <div class="modal" id="%s" aria-labelledby="%sLabel" aria-hidden="true">
          <div class="modal-background"></div>
          <div class="modal-content" id="%sLabel">
            <div class="box">
            <h2>%s</h2>
              <div class="bootstrap-table">
                  <div class="fixed-table-toolbar">
                      <div class="bars">
                        <div id="toolbar" class="padding-bottom-2">
                          <button id="add" type="button" class="waves-effect waves-light button">%s %s</button>
                          <button id="delete" type="button" class="waves-effect waves-light button" disabled>%s %s</button>
                        </div>
                      </div>
                  </div>
                  <div class="fixed-table-container">
                      <table id="multi-sort" class="table">
                          <thead>
                              <tr>
                                  <th></th>
                                  <th><div class="th-inner">%s</div></th>
                                  <th><div class="th-inner">%s</div></th>
                              </tr>
                          </thead>
                          <tbody></tbody>
                      </table>
                    </div>
                </div>
                <button type="button" class="waves-effect waves-light button" data-close>%s</button>
                <button type="button" class="waves-effect waves-light button multi-sort-order-button" data-close>%s</button>
            </div>
          </div>
        </div>
      `,multipleSortButton:'<button class="button multi-sort" data-target="%s" title="%s">%s</button>',multipleSortSelect:'<select class="%s %s browser-default">'}}}[$.fn.bootstrapTable.theme];$.extend($.fn.bootstrapTable.defaults.icons,bootstrap.icons),$.extend($.fn.bootstrapTable.defaults.html,bootstrap.html);const showSortModal=t=>{const o=t.sortModalSelector,d=`#${o}`,s=t.options;if(!$(d).hasClass("modal")){const r=Utils.sprintf(t.constants.html.multipleSortModal,o,o,o,t.options.formatMultipleSort(),Utils.sprintf(t.constants.html.icon,s.iconsPrefix,t.constants.icons.plus),t.options.formatAddLevel(),Utils.sprintf(t.constants.html.icon,s.iconsPrefix,t.constants.icons.minus),t.options.formatDeleteLevel(),t.options.formatColumn(),t.options.formatOrder(),t.options.formatCancel(),t.options.formatSort());$("body").append($(r)),t.$sortModal=$(d);const n=t.$sortModal.find("tbody > tr");if(t.$sortModal.off("click","#add").on("click","#add",()=>{const e=t.$sortModal.find(".multi-sort-name:first option").length;let l=t.$sortModal.find("tbody tr").length;l<e&&(l++,t.addLevel(),t.setButtonStates())}),t.$sortModal.off("click","#delete").on("click","#delete",()=>{const e=t.$sortModal.find(".multi-sort-name:first option").length;let l=t.$sortModal.find("tbody tr").length;l>1&&l<=e&&(l--,t.$sortModal.find("tbody tr:last").remove(),t.setButtonStates())}),t.$sortModal.off("click",".multi-sort-order-button").on("click",".multi-sort-order-button",()=>{const e=t.$sortModal.find("tbody > tr");let l=t.$sortModal.find("div.alert");const u=[],b=[],f=$.map(e,i=>{const a=$(i),p=a.find(".multi-sort-name").val(),m=a.find(".multi-sort-order").val();return u.push(p),{sortName:p,sortOrder:m}}),c=u.sort();for(let i=0;i<u.length-1;i++)c[i+1]===c[i]&&b.push(c[i]);b.length>0?l.length===0&&(l=`<div class="alert alert-danger" role="alert"><strong>${t.options.formatDuplicateAlertTitle()}</strong> ${t.options.formatDuplicateAlertDescription()}</div>`,$(l).insertBefore(t.$sortModal.find(".bars"))):(l.length===1&&$(l).remove(),$.inArray($.fn.bootstrapTable.theme,["bootstrap3","bootstrap4"])!==-1&&t.$sortModal.modal("hide"),t.multiSort(f))}),(t.options.sortPriority===null||t.options.sortPriority.length===0)&&t.options.sortName&&(t.options.sortPriority=[{sortName:t.options.sortName,sortOrder:t.options.sortOrder}]),t.options.sortPriority!==null&&t.options.sortPriority.length>0){if(n.length<t.options.sortPriority.length&&typeof t.options.sortPriority=="object")for(let e=0;e<t.options.sortPriority.length;e++)t.addLevel(e,t.options.sortPriority[e])}else t.addLevel(0);t.setButtonStates()}};$.fn.bootstrapTable.methods.push("multipleSort"),$.fn.bootstrapTable.methods.push("multiSort"),$.extend($.fn.bootstrapTable.defaults,{showMultiSort:!1,showMultiSortButton:!0,multiSortStrictSort:!1,sortPriority:null,onMultipleSort(){return!1}}),$.extend($.fn.bootstrapTable.Constructor.EVENTS,{"multiple-sort.bs.table":"onMultipleSort"}),$.extend($.fn.bootstrapTable.locales,{formatMultipleSort(){return"Multiple Sort"},formatAddLevel(){return"Add Level"},formatDeleteLevel(){return"Delete Level"},formatColumn(){return"Column"},formatOrder(){return"Order"},formatSortBy(){return"Sort by"},formatThenBy(){return"Then by"},formatSort(){return"Sort"},formatCancel(){return"Cancel"},formatDuplicateAlertTitle(){return"Duplicate(s) detected!"},formatDuplicateAlertDescription(){return"Please remove or change any duplicate column."},formatSortOrders(){return{asc:"Ascending",desc:"Descending"}}}),$.extend($.fn.bootstrapTable.defaults,$.fn.bootstrapTable.locales);const BootstrapTable=$.fn.bootstrapTable.Constructor,_initToolbar=BootstrapTable.prototype.initToolbar,_destroy=BootstrapTable.prototype.destroy;BootstrapTable.prototype.initToolbar=function(...t){this.showToolbar=this.showToolbar||this.options.showMultiSort;const o=this,d=`sortModal_${this.$el.attr("id")}`,s=`#${d}`;if(this.$sortModal=$(s),this.sortModalSelector=d,o.options.sortPriority!==null&&o.onMultipleSort(),_initToolbar.apply(this,Array.prototype.slice.apply(t)),o.options.sidePagination==="server"&&!isSingleSort&&o.options.sortPriority!==null){const r=o.options.queryParams;o.options.queryParams=n=>(n.multiSort=o.options.sortPriority,r(n))}if(this.options.showMultiSort){const r=this.$toolbar.find(">."+o.constants.classes.buttonsGroup.split(" ").join(".")).first();let n=this.$toolbar.find("div.multi-sort");const e=o.options;!n.length&&this.options.showMultiSortButton&&(n=Utils.sprintf(o.constants.html.multipleSortButton,o.sortModalSelector,this.options.formatMultipleSort(),Utils.sprintf(o.constants.html.icon,e.iconsPrefix,e.icons.sort)),r.append(n),$.fn.bootstrapTable.theme==="semantic"?this.$toolbar.find(".multi-sort").on("click",()=>{$(s).modal("show")}):$.fn.bootstrapTable.theme==="materialize"?this.$toolbar.find(".multi-sort").on("click",()=>{$(s).modal()}):$.fn.bootstrapTable.theme==="foundation"?this.$toolbar.find(".multi-sort").on("click",()=>{this.foundationModal||(this.foundationModal=new Foundation.Reveal($(s))),this.foundationModal.open()}):$.fn.bootstrapTable.theme==="bulma"&&this.$toolbar.find(".multi-sort").on("click",()=>{$("html").toggleClass("is-clipped"),$(s).toggleClass("is-active"),$("button[data-close]").one("click",()=>{$("html").toggleClass("is-clipped"),$(s).toggleClass("is-active")})}),showSortModal(o)),this.$el.on("sort.bs.table",()=>{isSingleSort=!0}),this.$el.on("multiple-sort.bs.table",()=>{isSingleSort=!1}),this.$el.on("load-success.bs.table",()=>{!isSingleSort&&o.options.sortPriority!==null&&typeof o.options.sortPriority=="object"&&o.options.sidePagination!=="server"&&o.onMultipleSort()}),this.$el.on("column-switch.bs.table",(l,u)=>{for(let b=0;b<o.options.sortPriority.length;b++)o.options.sortPriority[b].sortName===u&&o.options.sortPriority.splice(b,1);o.assignSortableArrows(),o.$sortModal.remove(),showSortModal(o)}),this.$el.on("reset-view.bs.table",()=>{!isSingleSort&&o.options.sortPriority!==null&&typeof o.options.sortPriority=="object"&&o.assignSortableArrows()})}},BootstrapTable.prototype.destroy=function(...t){_destroy.apply(this,Array.prototype.slice.apply(t)),this.options.showMultiSort&&this.$sortModal.remove()},BootstrapTable.prototype.multipleSort=function(){const t=this;!isSingleSort&&t.options.sortPriority!==null&&typeof t.options.sortPriority=="object"&&t.options.sidePagination!=="server"&&t.onMultipleSort()},BootstrapTable.prototype.onMultipleSort=function(){const t=this,o=(s,r)=>s>r?1:s<r?-1:0,d=(s,r)=>{const n=[],e=[];for(let l=0;l<t.options.sortPriority.length;l++){let u=t.options.sortPriority[l].sortName;const b=t.header.fields.indexOf(u),f=t.header.sorters[t.header.fields.indexOf(u)];t.header.sortNames[b]&&(u=t.header.sortNames[b]);const c=t.options.sortPriority[l].sortOrder==="desc"?-1:1;let i=Utils.getItemField(s,u),a=Utils.getItemField(r,u);const p=$.fn.bootstrapTable.utils.calculateObjectValue(t.header,f,[i,a]),m=$.fn.bootstrapTable.utils.calculateObjectValue(t.header,f,[a,i]);if(p!==void 0&&m!==void 0){n.push(c*p),e.push(c*m);continue}i==null&&(i=""),a==null&&(a=""),$.isNumeric(i)&&$.isNumeric(a)?(i=parseFloat(i),a=parseFloat(a)):(i=i.toString(),a=a.toString(),t.options.multiSortStrictSort&&(i=i.toLowerCase(),a=a.toLowerCase())),n.push(c*o(i,a)),e.push(c*o(a,i))}return o(n,e)};this.data.sort((s,r)=>d(s,r)),this.initBody(),this.assignSortableArrows(),this.trigger("multiple-sort")},BootstrapTable.prototype.addLevel=function(t,o){const d=t===0?this.options.formatSortBy():this.options.formatThenBy();this.$sortModal.find("tbody").append($("<tr>").append($("<td>").text(d)).append($("<td>").append($(Utils.sprintf(this.constants.html.multipleSortSelect,this.constants.classes.paginationDropdown,"multi-sort-name")))).append($("<td>").append($(Utils.sprintf(this.constants.html.multipleSortSelect,this.constants.classes.paginationDropdown,"multi-sort-order")))));const s=this.$sortModal.find(".multi-sort-name").last(),r=this.$sortModal.find(".multi-sort-order").last();$.each(this.columns,(n,e)=>{if(e.sortable===!1||e.visible===!1)return!0;s.append(`<option value="${e.field}">${e.title}</option>`)}),$.each(this.options.formatSortOrders(),(n,e)=>{r.append(`<option value="${n}">${e}</option>`)}),o!==void 0&&(s.find(`option[value="${o.sortName}"]`).attr("selected",!0),r.find(`option[value="${o.sortOrder}"]`).attr("selected",!0))},BootstrapTable.prototype.assignSortableArrows=function(){const t=this,o=t.$header.find("th");for(let d=0;d<o.length;d++)for(let s=0;s<t.options.sortPriority.length;s++)$(o[d]).data("field")===t.options.sortPriority[s].sortName&&$(o[d]).find(".sortable").removeClass("desc asc").addClass(t.options.sortPriority[s].sortOrder)},BootstrapTable.prototype.setButtonStates=function(){const t=this.$sortModal.find(".multi-sort-name:first option").length,o=this.$sortModal.find("tbody tr").length;o===t&&this.$sortModal.find("#add").attr("disabled","disabled"),o>1&&this.$sortModal.find("#delete").removeAttr("disabled"),o<t&&this.$sortModal.find("#add").removeAttr("disabled"),o===1&&this.$sortModal.find("#delete").attr("disabled","disabled")},BootstrapTable.prototype.multiSort=function(t){if(this.options.sortPriority=t,this.options.sortName="",this.options.sidePagination==="server"){this.options.queryParams=o=>(o.multiSort=this.options.sortPriority,$.fn.bootstrapTable.utils.calculateObjectValue(this.options,this.options.queryParams,[o])),isSingleSort=!1,this.initServer(this.options.silentSort);return}this.onMultipleSort()};
