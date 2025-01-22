const Utils=$.fn.bootstrapTable.utils;function printPageBuilderDefault(s){return`
  <html>
  <head>
  <style type="text/css" media="print">
  @page {
    size: auto;
    margin: 25px 0 25px 0;
  }
  </style>
  <style type="text/css" media="all">
  table {
    border-collapse: collapse;
    font-size: 12px;
  }
  table, th, td {
    border: 1px solid grey;
  }
  th, td {
    text-align: center;
    vertical-align: middle;
  }
  p {
    font-weight: bold;
    margin-left:20px;
  }
  table {
    width:94%;
    margin-left:3%;
    margin-right:3%;
  }
  div.bs-table-print {
    text-align:center;
  }
  </style>
  </head>
  <title>Print Table</title>
  <body>
  <p>Printed on: ${new Date} </p>
  <div class="bs-table-print">${s}</div>
  </body>
  </html>`}$.extend($.fn.bootstrapTable.defaults,{showPrint:!1,printAsFilteredAndSortedOnUI:!0,printSortColumn:void 0,printSortOrder:"asc",printPageBuilder(s){return printPageBuilderDefault(s)}}),$.extend($.fn.bootstrapTable.COLUMN_DEFAULTS,{printFilter:void 0,printIgnore:!1,printFormatter:void 0}),$.extend($.fn.bootstrapTable.defaults.icons,{print:{bootstrap3:"glyphicon-print icon-share"}[$.fn.bootstrapTable.theme]||"fa-print"}),$.BootstrapTable=class extends $.BootstrapTable{initToolbar(...s){if(this.showToolbar=this.showToolbar||this.options.showPrint,super.initToolbar(...s),!this.options.showPrint)return;const d=this.$toolbar.find(">.columns");let l=d.find("button.bs-print");l.length||(l=$(`
        <button class="${this.constants.buttonsClass} bs-print" type="button">
        <i class="${this.options.iconsPrefix} ${this.options.icons.print}"></i>
        </button>`).appendTo(d)),l.off("click").on("click",()=>{this.doPrint(this.options.printAsFilteredAndSortedOnUI?this.getData():this.options.data.slice(0))})}doPrint(s){const d=(e,t,i)=>{const n=Utils.calculateObjectValue(i,i.printFormatter,[e[i.field],e,t],e[i.field]);return typeof n>"u"||n===null?this.options.undefinedText:n},l=(e,t)=>{const n=[`<table dir="${this.$el.attr("dir")||"ltr"}"><thead>`];for(const r of t){n.push("<tr>");for(let o=0;o<r.length;o++)r[o].printIgnore||n.push(`<th
              ${Utils.sprintf(' rowspan="%s"',r[o].rowspan)}
              ${Utils.sprintf(' colspan="%s"',r[o].colspan)}
              >${r[o].title}</th>`);n.push("</tr>")}n.push("</thead><tbody>");for(let r=0;r<e.length;r++){n.push("<tr>");for(const o of t)for(let a=0;a<o.length;a++)!o[a].printIgnore&&o[a].field&&n.push("<td>",d(e[r],r,o[a]),"</td>");n.push("</tr>")}return n.push("</tbody></table>"),n.join("")},h=(e,t,i)=>{if(!t)return e;let n=i!=="asc";return n=-(+n||-1),e.sort((r,o)=>n*r[t].localeCompare(o[t]))},u=(e,t)=>{for(let i=0;i<t.length;++i)if(e[t[i].colName]!==t[i].value)return!1;return!0};s=((e,t)=>e.filter(i=>u(i,t)))(s,(e=>!e||!e[0]?[]:e[0].filter(t=>t.printFilter).map(t=>({colName:t.field,value:t.printFilter})))(this.options.columns)),s=h(s,this.options.printSortColumn,this.options.printSortOrder);const f=l(s,this.options.columns),p=window.open("");p.document.write(this.options.printPageBuilder.call(this,f)),p.document.close(),p.focus(),p.print(),p.close()}};
