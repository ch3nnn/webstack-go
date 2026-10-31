import{aD as $e,h as w,o as ze,P as L,H as ne,K as ke,k as M,u as l,v as h,z as b,V as R,f as _e,aE as W,j as Re,aF as me,aw as ve,as as ie,U as se,B as S,a7 as Pe,W as J,Z as U,_ as Be,aG as Ne,aH as fe,aa as De,E as pe,m as d,y,a6 as B,aI as le,aJ as ae,aK as ce,aL as de,aM as Ge,I as ue,p as T,n as x,N as E,R as A,aN as K,a1 as ye,a8 as O,l as Z,q as ee,aO as Me,x as P,ai as G,ak as z,ao as Ee,ap as Ie,aP as We,aQ as X,ax as I,aR as Ve}from"./index-Bpd5eKPv.js";import{f as j,_ as xe}from"./_plugin-vue_export-helper-CtyYQGco.js";let Q=[];const be=new WeakMap;function Le(){Q.forEach(e=>e(...be.get(e))),Q=[]}function Oe(e,...t){be.set(e,t),!Q.includes(e)&&Q.push(e)===1&&requestAnimationFrame(Le)}function je(e){if(typeof e=="number")return{"":e.toString()};const t={};return e.split(/ +/).forEach(o=>{if(o==="")return;const[a,i]=o.split(":");i===void 0?t[""]=a:t[a]=i}),t}function F(e,t){var o;if(e==null)return;const a=je(e);if(t===void 0)return a[""];if(typeof t=="string")return(o=a[t])!==null&&o!==void 0?o:a[""];if(Array.isArray(t)){for(let i=t.length-1;i>=0;--i){const r=t[i];if(r in a)return a[r]}return a[""]}else{let i,r=-1;return Object.keys(a).forEach(n=>{const s=Number(n);!Number.isNaN(s)&&t>=s&&s>=r&&(r=s,i=a[n])}),i}}const Fe={xs:0,s:640,m:1024,l:1280,xl:1536,"2xl":1920};function Te(e){return`(min-width: ${e}px)`}const H={};function Ae(e=Fe){if(!$e)return w(()=>[]);if(typeof window.matchMedia!="function")return w(()=>[]);const t=L({}),o=Object.keys(e),a=(i,r)=>{i.matches?t.value[r]=!0:t.value[r]=!1};return o.forEach(i=>{const r=e[i];let n,s;H[r]===void 0?(n=window.matchMedia(Te(r)),n.addEventListener?n.addEventListener("change",c=>{s.forEach(g=>{g(c,i)})}):n.addListener&&n.addListener(c=>{s.forEach(g=>{g(c,i)})}),s=new Set,H[r]={mql:n,cbs:s}):(n=H[r].mql,s=H[r].cbs),s.add(a),n.matches&&s.forEach(c=>{c(n,i)})}),ze(()=>{o.forEach(i=>{const{cbs:r}=H[e[i]];r.has(a)&&r.delete(a)})}),w(()=>{const{value:i}=t;return o.filter(r=>i[r])})}function qe(e,t="default",o=[]){const a=e.$slots[t];return a===void 0?o:a()}function He(e){const{infoColor:t,successColor:o,warningColor:a,errorColor:i,textColor2:r,progressRailColor:n,fontSize:s,fontWeight:c}=e;return{fontSize:s,fontSizeCircle:"28px",fontWeightCircle:c,railColor:n,railHeight:"8px",iconSizeCircle:"36px",iconSizeLine:"18px",iconColor:t,iconColorInfo:t,iconColorSuccess:o,iconColorWarning:a,iconColorError:i,textColorCircle:r,textColorLineInner:"rgb(255, 255, 255)",textColorLineOuter:r,fillColor:t,fillColorInfo:t,fillColorSuccess:o,fillColorWarning:a,fillColorError:i,lineBgProcessing:"linear-gradient(90deg, rgba(255, 255, 255, .3) 0%, rgba(255, 255, 255, .5) 100%)"}}const Ue={common:ne,self:He};var Xe={titleFontSizeSmall:"26px",titleFontSizeMedium:"32px",titleFontSizeLarge:"40px",titleFontSizeHuge:"48px",fontSizeSmall:"14px",fontSizeMedium:"14px",fontSizeLarge:"15px",fontSizeHuge:"16px",iconSizeSmall:"64px",iconSizeMedium:"80px",iconSizeLarge:"100px",iconSizeHuge:"125px",iconColor418:void 0,iconColor404:void 0,iconColor403:void 0,iconColor500:void 0};function Ye(e){const{textColor2:t,textColor1:o,errorColor:a,successColor:i,infoColor:r,warningColor:n,lineHeight:s,fontWeightStrong:c}=e;return{...Xe,lineHeight:s,titleFontWeight:c,titleTextColor:o,textColor:t,iconColorError:a,iconColorSuccess:i,iconColorInfo:r,iconColorWarning:n}}const Qe={common:ne,self:Ye},Ce=ke("n-grid"),Je={span:{type:[Number,String],default:1},offset:{type:[Number,String],default:0},suffix:Boolean,privateOffset:Number,privateSpan:Number,privateColStart:Number,privateShow:{type:Boolean,default:!0}};var te=M({__GRID_ITEM__:!0,name:"GridItem",alias:["Gi"],props:Je,setup(){const{isSsrRef:e,xGapRef:t,itemStyleRef:o,overflowRef:a,layoutShiftDisabledRef:i}=_e(Ce),r=Re();return{overflow:a,itemStyle:o,layoutShiftDisabled:i,mergedXGap:w(()=>W(t.value||0)),deriveStyle:()=>{e.value;const{privateSpan:n=1,privateShow:s=!0,privateColStart:c=void 0,privateOffset:g=0}=r.vnode.props,{value:v}=t,u=W(v||0);return{display:s?"":"none",gridColumn:`${c??`span ${n}`} / span ${n}`,marginLeft:g?`calc((100% - (${n} - 1) * ${u}) / ${n} * ${g} + ${u} * ${g})`:""}}}},render(){if(this.layoutShiftDisabled){const{span:e,offset:t,mergedXGap:o}=this;return l(),h("div",{key:1,style:R({gridColumn:`span ${e} / span ${e}`,marginLeft:t?`calc((100% - (${e} - 1) * ${o}) / ${e} * ${t} + ${o} * ${t})`:""})},[b(()=>this.$slots.default?.())],4)}return l(),h("div",{style:R([this.itemStyle,this.deriveStyle()])},[b(()=>this.$slots.default?.({overflow:this.overflow}))],4)}});let ge=!1;function Ze(){if(me&&window.CSS&&!ge&&(ge=!0,"registerProperty"in window?.CSS))try{CSS.registerProperty({name:"--n-color-start",syntax:"<color>",inherits:!1,initialValue:"#0000"}),CSS.registerProperty({name:"--n-color-end",syntax:"<color>",inherits:!1,initialValue:"#0000"})}catch{}}function Ke(e){const t=e.dirs?.find(({dir:o})=>o===ve);return!!(t&&t.value===!1)}const et={xs:0,s:640,m:1024,l:1280,xl:1536,xxl:1920},Se=24,re="__ssr__",tt={layoutShiftDisabled:Boolean,responsive:{type:[String,Boolean],default:"self"},cols:{type:[Number,String],default:Se},itemResponsive:Boolean,collapsed:Boolean,collapsedRows:{type:Number,default:1},itemStyle:[Object,String],xGap:{type:[Number,String],default:0},yGap:{type:[Number,String],default:0}};var rt=M({name:"Grid",inheritAttrs:!1,props:tt,setup(e){const{mergedClsPrefixRef:t,mergedBreakpointsRef:o}=J(e),a=/^\d+$/,i=L(void 0),r=Ae(o?.value||et),n=U(()=>!!(e.itemResponsive||!a.test(e.cols.toString())||!a.test(e.xGap.toString())||!a.test(e.yGap.toString()))),s=w(()=>{if(n.value)return e.responsive==="self"?i.value:r.value}),c=U(()=>Number(F(e.cols.toString(),s.value))??Se),g=U(()=>F(e.xGap.toString(),s.value)),v=U(()=>F(e.yGap.toString(),s.value)),u=$=>{i.value=$.contentRect.width},m=$=>{Oe(u,$)},C=L(!1),p=w(()=>{if(e.responsive==="self")return m}),f=L(!1),k=L();return Be(()=>{const{value:$}=k;$&&$.hasAttribute(re)&&($.removeAttribute(re),f.value=!0)}),De(Ce,{layoutShiftDisabledRef:pe(e,"layoutShiftDisabled"),isSsrRef:f,itemStyleRef:pe(e,"itemStyle"),xGapRef:g,overflowRef:C}),{isSsr:!me,contentEl:k,mergedClsPrefix:t,style:w(()=>e.layoutShiftDisabled?{width:"100%",display:"grid",gridTemplateColumns:`repeat(${e.cols}, minmax(0, 1fr))`,columnGap:W(e.xGap),rowGap:W(e.yGap)}:{width:"100%",display:"grid",gridTemplateColumns:`repeat(${c.value}, minmax(0, 1fr))`,columnGap:W(g.value),rowGap:W(v.value)}),isResponsive:n,responsiveQuery:s,responsiveCols:c,handleResize:p,overflow:C}},render(){if(this.layoutShiftDisabled)return ie("div",se({ref:"contentEl",class:`${this.mergedClsPrefix}-grid`,style:this.style},this.$attrs),this.$slots);const e=()=>{this.overflow=!1;const t=Ne(qe(this)),o=[],{collapsed:a,collapsedRows:i,responsiveCols:r,responsiveQuery:n}=this;t.forEach(u=>{if(u?.type?.__GRID_ITEM__!==!0)return;if(Ke(u)){const p=fe(u);p.props?p.props.privateShow=!1:p.props={privateShow:!1},o.push({child:p,rawChildSpan:0});return}u.dirs=u.dirs?.filter(({dir:p})=>p!==ve)||null,u.dirs?.length===0&&(u.dirs=null);const m=fe(u),C=Number(F(m.props?.span,n)??1);C!==0&&o.push({child:m,rawChildSpan:C})});let s=0;const c=o[o.length-1]?.child;if(c?.props){const u=c.props?.suffix;u!==void 0&&u!==!1&&(s=Number(F(c.props?.span,n)??1),c.props.privateSpan=s,c.props.privateColStart=r+1-s,c.props.privateShow=c.props.privateShow??!0)}let g=0,v=!1;for(const{child:u,rawChildSpan:m}of o){if(v&&(this.overflow=!0),!v){const C=Number(F(u.props?.offset,n)??0),p=Math.min(m+C,r);if(u.props?(u.props.privateSpan=p,u.props.privateOffset=C):u.props={privateSpan:p,privateOffset:C},a){const f=g%r;p+f>r&&(g+=r-f),p+g+s>i*r?v=!0:g+=p}}v&&(u.props?u.props.privateShow!==!0&&(u.props.privateShow=!1):u.props={privateShow:!1})}return ie("div",se({ref:"contentEl",class:`${this.mergedClsPrefix}-grid`,style:this.style,[re]:this.isSsr||void 0},this.$attrs),o.map(({child:u})=>u))};return this.isResponsive&&this.responsive==="self"?(l(),S(Pe,{key:1,onResize:this.handleResize},{default:e},1032,["onResize"])):e()}});const ot=["id"],it=["stop-color"],st=["stop-color"],nt=["viewBox"],lt=["d","stroke-width"],at=["d","stroke-width"],ct={success:(l(),S(de)),error:(l(),S(ce)),warning:(l(),S(ae)),info:(l(),S(le))};var dt=M({name:"ProgressCircle",props:{clsPrefix:{type:String,required:!0},status:{type:String,required:!0},strokeWidth:{type:Number,required:!0},fillColor:[String,Object],railColor:String,railStyle:[String,Object],percentage:{type:Number,default:0},offsetDegree:{type:Number,default:0},showIndicator:{type:Boolean,required:!0},indicatorTextColor:String,unit:String,viewBoxWidth:{type:Number,required:!0},gapDegree:{type:Number,required:!0},gapOffsetDegree:{type:Number,default:0}},setup(e,{slots:t}){const o=w(()=>{const r="gradient",{fillColor:n}=e;return typeof n=="object"?`${r}-${Ge(JSON.stringify(n))}`:r});function a(r,n,s,c){const{gapDegree:g,viewBoxWidth:v,strokeWidth:u}=e,m=50,C=0,p=m,f=0,k=100,$=50+u/2,_=`M ${$},${$} m ${C},${p}
      a ${m},${m} 0 1 1 ${f},-100
      a ${m},${m} 0 1 1 0,${k}`,N=Math.PI*2*m;return{pathString:_,pathStyle:{stroke:c==="rail"?s:typeof e.fillColor=="object"?`url(#${o.value})`:s,strokeDasharray:`${Math.min(r,100)/100*(N-g)}px ${v*8}px`,strokeDashoffset:`-${g/2}px`,transformOrigin:n?"center":void 0,transform:n?`rotate(${n}deg)`:void 0}}}const i=()=>{const r=typeof e.fillColor=="object",n=r?e.fillColor.stops[0]:"",s=r?e.fillColor.stops[1]:"";return r&&(l(),h("defs",null,[d("linearGradient",{id:o.value,x1:"0%",y1:"100%",x2:"100%",y2:"0%"},[d("stop",{offset:"0%","stop-color":n},null,8,it),d("stop",{offset:"100%","stop-color":s},null,8,st)],8,ot)]))};return()=>{const{fillColor:r,railColor:n,strokeWidth:s,offsetDegree:c,status:g,percentage:v,showIndicator:u,indicatorTextColor:m,unit:C,gapOffsetDegree:p,clsPrefix:f}=e,{pathString:k,pathStyle:$}=a(100,0,n,"rail"),{pathString:_,pathStyle:N}=a(v,c,r,"fill"),D=100+s;return l(),h("div",{class:y(`${f}-progress-content`),role:"none"},[d("div",{class:y(`${f}-progress-graph`),"aria-hidden":!0},[d("div",{class:y(`${f}-progress-graph-circle`),style:R({transform:p?`rotate(${p}deg)`:void 0})},[(l(),h("svg",{viewBox:`0 0 ${D} ${D}`},[b(()=>i()),d("g",null,[d("path",{class:y(`${f}-progress-graph-circle-rail`),d:k,"stroke-width":s,"stroke-linecap":"round",fill:"none",style:R($)},null,14,lt)]),d("g",null,[d("path",{class:y([`${f}-progress-graph-circle-fill`,v===0&&`${f}-progress-graph-circle-fill--empty`]),d:_,"stroke-width":s,"stroke-linecap":"round",fill:"none",style:R(N)},null,14,at)])],8,nt))],6)],2),u?(l(),h("div",{key:0},[t.default?(l(),h("div",{key:0,class:y(`${f}-progress-custom-content`),role:"none"},[b(()=>t.default())],2)):(l(),h(B,{key:1},[g!=="default"?(l(),h("div",{key:0,class:y(`${f}-progress-icon`),"aria-hidden":!0},[(l(),S(ue,{clsPrefix:f},{default:()=>ct[g]},1032,["clsPrefix"]))],2)):(l(),h("div",{key:1,class:y(`${f}-progress-text`),style:R({color:m}),role:"none"},[d("span",{class:y(`${f}-progress-text__percentage`)},[b(()=>v)],2),d("span",{class:y(`${f}-progress-text__unit`)},[b(()=>C)],2)],6))],64))])):b(()=>null)],2)}}});const ut={success:(l(),S(de)),error:(l(),S(ce)),warning:(l(),S(ae)),info:(l(),S(le))};var ft=M({name:"ProgressLine",props:{clsPrefix:{type:String,required:!0},percentage:{type:Number,default:0},railColor:String,railStyle:[String,Object],fillColor:[String,Object],status:{type:String,required:!0},indicatorPlacement:{type:String,required:!0},indicatorTextColor:String,unit:{type:String,default:"%"},processing:{type:Boolean,required:!0},showIndicator:{type:Boolean,required:!0},height:[String,Number],railBorderRadius:[String,Number],fillBorderRadius:[String,Number]},setup(e,{slots:t}){const o=w(()=>j(e.height)),a=w(()=>typeof e.fillColor=="object"?`linear-gradient(to right, ${e.fillColor?.stops[0]} , ${e.fillColor?.stops[1]})`:e.fillColor),i=w(()=>e.railBorderRadius!==void 0?j(e.railBorderRadius):e.height!==void 0?j(e.height,{c:.5}):""),r=w(()=>e.fillBorderRadius!==void 0?j(e.fillBorderRadius):e.railBorderRadius!==void 0?j(e.railBorderRadius):e.height!==void 0?j(e.height,{c:.5}):"");return()=>{const{indicatorPlacement:n,railColor:s,railStyle:c,percentage:g,unit:v,indicatorTextColor:u,status:m,showIndicator:C,processing:p,clsPrefix:f}=e;return l(),h("div",{class:y(`${f}-progress-content`),role:"none"},[d("div",{class:y(`${f}-progress-graph`),"aria-hidden":!0},[d("div",{class:y([`${f}-progress-graph-line`,{[`${f}-progress-graph-line--indicator-${n}`]:!0}])},[d("div",{class:y(`${f}-progress-graph-line-rail`),style:R([{backgroundColor:s,height:o.value,borderRadius:i.value},c])},[d("div",{class:y([`${f}-progress-graph-line-fill`,p&&`${f}-progress-graph-line-fill--processing`]),style:R({maxWidth:`${e.percentage}%`,background:a.value,height:o.value,lineHeight:o.value,borderRadius:r.value})},[n==="inside"?(l(),h("div",{key:0,class:y(`${f}-progress-graph-line-indicator`),style:R({color:u})},[t.default?(l(),h(B,{key:0},[b(()=>t.default())],64)):(l(),h(B,{key:1},[b(()=>`${g}${v}`)],64))],6)):b(()=>null)],6)],6)],2)],2),C&&n==="outside"?(l(),h("div",{key:0},[t.default?(l(),h("div",{key:0,class:y(`${f}-progress-custom-content`),style:R({color:u}),role:"none"},[b(()=>t.default())],6)):(l(),h(B,{key:1},[m==="default"?(l(),h("div",{key:0,role:"none",class:y(`${f}-progress-icon ${f}-progress-icon--as-text`),style:R({color:u})},[b(()=>g),b(()=>v)],6)):(l(),h("div",{key:1,class:y(`${f}-progress-icon`),"aria-hidden":!0},[(l(),S(ue,{clsPrefix:f},{default:()=>ut[m]},1032,["clsPrefix"]))],2))],64))])):b(()=>null)],2)}}});const pt=["id"],gt=["stop-color"],ht=["stop-color"],mt=["d","stroke-width"],vt=["d","stroke-width"],yt=["viewBox"];function he(e,t,o=100){return`m ${o/2} ${o/2-e} a ${e} ${e} 0 1 1 0 ${2*e} a ${e} ${e} 0 1 1 0 -${2*e}`}var xt=M({name:"ProgressMultipleCircle",props:{clsPrefix:{type:String,required:!0},viewBoxWidth:{type:Number,required:!0},percentage:{type:Array,default:[0]},strokeWidth:{type:Number,required:!0},circleGap:{type:Number,required:!0},showIndicator:{type:Boolean,required:!0},fillColor:{type:Array,default:()=>[]},railColor:{type:Array,default:()=>[]},railStyle:{type:Array,default:()=>[]}},setup(e,{slots:t}){const o=w(()=>e.percentage.map((i,r)=>`${Math.PI*i/100*(e.viewBoxWidth/2-e.strokeWidth/2*(1+2*r)-e.circleGap*r)*2}, ${e.viewBoxWidth*8}`)),a=(i,r)=>{const n=e.fillColor[r],s=typeof n=="object"?n.stops[0]:"",c=typeof n=="object"?n.stops[1]:"";return typeof e.fillColor[r]=="object"&&(l(),h("linearGradient",{id:`gradient-${r}`,x1:"100%",y1:"0%",x2:"0%",y2:"100%"},[d("stop",{offset:"0%","stop-color":s},null,8,gt),d("stop",{offset:"100%","stop-color":c},null,8,ht)],8,pt))};return()=>{const{viewBoxWidth:i,strokeWidth:r,circleGap:n,showIndicator:s,fillColor:c,railColor:g,railStyle:v,percentage:u,clsPrefix:m}=e;return l(),h("div",{class:y(`${m}-progress-content`),role:"none"},[d("div",{class:y(`${m}-progress-graph`),"aria-hidden":!0},[d("div",{class:y(`${m}-progress-graph-circle`)},[(l(),h("svg",{viewBox:`0 0 ${i} ${i}`},[d("defs",null,[b(()=>u.map((C,p)=>a(C,p)))]),b(()=>u.map((C,p)=>(l(),h("g",{key:p},[d("path",{class:y(`${m}-progress-graph-circle-rail`),d:he(i/2-r/2*(1+2*p)-n*p,r,i),"stroke-width":r,"stroke-linecap":"round",fill:"none",style:R([{strokeDashoffset:0,stroke:g[p]},v[p]])},null,14,mt),d("path",{class:y([`${m}-progress-graph-circle-fill`,C===0&&`${m}-progress-graph-circle-fill--empty`]),d:he(i/2-r/2*(1+2*p)-n*p,r,i),"stroke-width":r,"stroke-linecap":"round",fill:"none",style:R({strokeDasharray:o.value[p],strokeDashoffset:0,stroke:typeof c[p]=="object"?`url(#gradient-${p})`:c[p]})},null,14,vt)]))))],8,yt))],2)],2),s&&t.default?(l(),h("div",{key:0},[d("div",{class:y(`${m}-progress-text`)},[b(()=>t.default())],2)])):b(()=>null)],2)}}}),bt=T([x("progress",{display:"inline-block"},[x("progress-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 `),E("line",`
 width: 100%;
 display: block;
 `,[x("progress-content",`
 display: flex;
 align-items: center;
 `,[x("progress-graph",{flex:1})]),x("progress-custom-content",{marginLeft:"14px"}),x("progress-icon",`
 width: 30px;
 padding-left: 14px;
 height: var(--n-icon-size-line);
 line-height: var(--n-icon-size-line);
 font-size: var(--n-icon-size-line);
 `,[E("as-text",`
 color: var(--n-text-color-line-outer);
 text-align: center;
 width: 40px;
 font-size: var(--n-font-size);
 padding-left: 4px;
 transition: color .3s var(--n-bezier);
 `)])]),E("circle, dashboard",{width:"120px"},[x("progress-custom-content",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 display: flex;
 align-items: center;
 justify-content: center;
 `),x("progress-text",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 display: flex;
 align-items: center;
 color: inherit;
 font-size: var(--n-font-size-circle);
 color: var(--n-text-color-circle);
 font-weight: var(--n-font-weight-circle);
 transition: color .3s var(--n-bezier);
 white-space: nowrap;
 `),x("progress-icon",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 display: flex;
 align-items: center;
 color: var(--n-icon-color);
 font-size: var(--n-icon-size-circle);
 `)]),E("multiple-circle",`
 width: 200px;
 color: inherit;
 `,[x("progress-text",`
 font-weight: var(--n-font-weight-circle);
 color: var(--n-text-color-circle);
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 display: flex;
 align-items: center;
 justify-content: center;
 transition: color .3s var(--n-bezier);
 `)]),x("progress-content",{position:"relative"}),x("progress-graph",{position:"relative"},[x("progress-graph-circle",[T("svg",{verticalAlign:"bottom"}),x("progress-graph-circle-fill",`
 stroke: var(--n-fill-color);
 transition:
 opacity .3s var(--n-bezier),
 stroke .3s var(--n-bezier),
 stroke-dasharray .3s var(--n-bezier);
 `,[E("empty",{opacity:0})]),x("progress-graph-circle-rail",`
 transition: stroke .3s var(--n-bezier);
 overflow: hidden;
 stroke: var(--n-rail-color);
 `)]),x("progress-graph-line",[E("indicator-inside",[x("progress-graph-line-rail",`
 height: 16px;
 line-height: 16px;
 border-radius: 10px;
 `,[x("progress-graph-line-fill",`
 height: inherit;
 border-radius: 10px;
 `),x("progress-graph-line-indicator",`
 background: #0000;
 white-space: nowrap;
 text-align: right;
 margin-left: 14px;
 margin-right: 14px;
 height: inherit;
 font-size: 12px;
 color: var(--n-text-color-line-inner);
 transition: color .3s var(--n-bezier);
 `)])]),E("indicator-inside-label",`
 height: 16px;
 display: flex;
 align-items: center;
 `,[x("progress-graph-line-rail",`
 flex: 1;
 transition: background-color .3s var(--n-bezier);
 `),x("progress-graph-line-indicator",`
 background: var(--n-fill-color);
 font-size: 12px;
 transform: translateZ(0);
 display: flex;
 vertical-align: middle;
 height: 16px;
 line-height: 16px;
 padding: 0 10px;
 border-radius: 10px;
 position: absolute;
 white-space: nowrap;
 color: var(--n-text-color-line-inner);
 transition:
 right .2s var(--n-bezier),
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `)]),x("progress-graph-line-rail",`
 position: relative;
 overflow: hidden;
 height: var(--n-rail-height);
 border-radius: 5px;
 background-color: var(--n-rail-color);
 transition: background-color .3s var(--n-bezier);
 `,[x("progress-graph-line-fill",`
 background: var(--n-fill-color);
 position: relative;
 border-radius: 5px;
 height: inherit;
 width: 100%;
 max-width: 0%;
 transition:
 background-color .3s var(--n-bezier),
 max-width .2s var(--n-bezier);
 `,[E("processing",[T("&::after",`
 content: "";
 background-image: var(--n-line-bg-processing);
 animation: progress-processing-animation 2s var(--n-bezier) infinite;
 `)])])])])])]),T("@keyframes progress-processing-animation",`
 0% {
 position: absolute;
 left: 0;
 top: 0;
 bottom: 0;
 right: 100%;
 opacity: 1;
 }
 66% {
 position: absolute;
 left: 0;
 top: 0;
 bottom: 0;
 right: 0;
 opacity: 0;
 }
 100% {
 position: absolute;
 left: 0;
 top: 0;
 bottom: 0;
 right: 0;
 opacity: 0;
 }
 `)]);const Ct=["aria-valuenow","role"],St={...A.props,processing:Boolean,type:{type:String,default:"line"},gapDegree:Number,gapOffsetDegree:Number,status:{type:String,default:"default"},railColor:[String,Array],railStyle:[String,Array],color:[String,Array,Object],viewBoxWidth:{type:Number,default:100},strokeWidth:{type:Number,default:7},percentage:[Number,Array],unit:{type:String,default:"%"},showIndicator:{type:Boolean,default:!0},indicatorPosition:{type:String,default:"outside"},indicatorPlacement:{type:String,default:"outside"},indicatorTextColor:String,circleGap:{type:Number,default:1},height:Number,borderRadius:[String,Number],fillBorderRadius:[String,Number],offsetDegree:Number};var oe=M({name:"Progress",props:St,setup(e){const t=w(()=>e.indicatorPlacement||e.indicatorPosition),o=w(()=>{if(e.gapDegree||e.gapDegree===0)return e.gapDegree;if(e.type==="dashboard")return 75}),{mergedClsPrefixRef:a,inlineThemeDisabled:i}=J(e),r=A("Progress","-progress",bt,Ue,e,a),n=w(()=>{const{status:c}=e,{common:{cubicBezierEaseInOut:g},self:{fontSize:v,fontSizeCircle:u,railColor:m,railHeight:C,iconSizeCircle:p,iconSizeLine:f,textColorCircle:k,textColorLineInner:$,textColorLineOuter:_,lineBgProcessing:N,fontWeightCircle:D,[O("iconColor",c)]:V,[O("fillColor",c)]:q}}=r.value;return{"--n-bezier":g,"--n-fill-color":q,"--n-font-size":v,"--n-font-size-circle":u,"--n-font-weight-circle":D,"--n-icon-color":V,"--n-icon-size-circle":p,"--n-icon-size-line":f,"--n-line-bg-processing":N,"--n-rail-color":m,"--n-rail-height":C,"--n-text-color-circle":k,"--n-text-color-line-inner":$,"--n-text-color-line-outer":_}}),s=i?ye("progress",w(()=>e.status[0]),n,e):void 0;return{mergedClsPrefix:a,mergedIndicatorPlacement:t,gapDeg:o,cssVars:i?void 0:n,themeClass:s?.themeClass,onRender:s?.onRender}},render(){const{type:e,cssVars:t,indicatorTextColor:o,showIndicator:a,status:i,railColor:r,railStyle:n,color:s,percentage:c,viewBoxWidth:g,strokeWidth:v,mergedIndicatorPlacement:u,unit:m,borderRadius:C,fillBorderRadius:p,height:f,processing:k,circleGap:$,mergedClsPrefix:_,gapDeg:N,gapOffsetDegree:D,themeClass:V,$slots:q,onRender:we}=this;return we?.(),l(),h("div",{class:y([V,`${_}-progress`,`${_}-progress--${e}`,`${_}-progress--${i}`]),style:R(t),"aria-valuemax":100,"aria-valuemin":0,"aria-valuenow":c,role:e==="circle"||e==="line"||e==="dashboard"?"progressbar":"none"},[e==="circle"||e==="dashboard"?(l(),S(dt,{key:0,clsPrefix:_,status:i,showIndicator:a,indicatorTextColor:o,railColor:r,fillColor:s,railStyle:n,offsetDegree:this.offsetDegree,percentage:c,viewBoxWidth:g,strokeWidth:v,gapDegree:N===void 0?e==="dashboard"?75:0:N,gapOffsetDegree:D,unit:m},K(q),1032,["clsPrefix","status","showIndicator","indicatorTextColor","railColor","fillColor","railStyle","offsetDegree","percentage","viewBoxWidth","strokeWidth","gapDegree","gapOffsetDegree","unit"])):(l(),h(B,{key:1},[e==="line"?(l(),S(ft,{key:0,clsPrefix:_,status:i,showIndicator:a,indicatorTextColor:o,railColor:r,fillColor:s,railStyle:n,percentage:c,processing:k,indicatorPlacement:u,unit:m,fillBorderRadius:p,railBorderRadius:C,height:f},K(q),1032,["clsPrefix","status","showIndicator","indicatorTextColor","railColor","fillColor","railStyle","percentage","processing","indicatorPlacement","unit","fillBorderRadius","railBorderRadius","height"])):(l(),h(B,{key:1},[e==="multiple-circle"?(l(),S(xt,{key:0,clsPrefix:_,strokeWidth:v,railColor:r,fillColor:s,railStyle:n,viewBoxWidth:g,percentage:c,showIndicator:a,circleGap:$},K(q),1032,["clsPrefix","strokeWidth","railColor","fillColor","railStyle","viewBoxWidth","percentage","showIndicator","circleGap"])):b(()=>null)],64))],64))],14,Ct)}});function wt(){return(()=>{const e=Z("c69907197a2424a9");return e[0]||(e[0]=d("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 36 36"},[d("path",{fill:"#EF9645",d:"M15.5 2.965c1.381 0 2.5 1.119 2.5 2.5v.005L20.5.465c1.381 0 2.5 1.119 2.5 2.5V4.25l2.5-1.535c1.381 0 2.5 1.119 2.5 2.5V8.75L29 18H15.458L15.5 2.965z"}),d("path",{fill:"#FFDC5D",d:"M4.625 16.219c1.381-.611 3.354.208 4.75 2.188.917 1.3 1.187 3.151 2.391 3.344.46.073 1.234-.313 1.234-1.397V4.5s0-2 2-2 2 2 2 2v11.633c0-.029 1-.064 1-.082V2s0-2 2-2 2 2 2 2v14.053c0 .017 1 .041 1 .069V4.25s0-2 2-2 2 2 2 2v12.638c0 .118 1 .251 1 .398V8.75s0-2 2-2 2 2 2 2V24c0 6.627-5.373 12-12 12-4.775 0-8.06-2.598-9.896-5.292C8.547 28.423 8.096 26.051 8 25.334c0 0-.123-1.479-1.156-2.865-1.469-1.969-2.5-3.156-3.125-3.866-.317-.359-.625-1.707.906-2.384z"})],-1))})()}function $t(){return(()=>{const e=Z("3c80ed8b3f0240f8");return e[0]||(e[0]=d("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 36 36"},[d("circle",{fill:"#FFCB4C",cx:"18",cy:"17.018",r:"17"}),d("path",{fill:"#65471B",d:"M14.524 21.036c-.145-.116-.258-.274-.312-.464-.134-.46.13-.918.59-1.021 4.528-1.021 7.577 1.363 7.706 1.465.384.306.459.845.173 1.205-.286.358-.828.401-1.211.097-.11-.084-2.523-1.923-6.182-1.098-.274.061-.554-.016-.764-.184z"}),d("ellipse",{fill:"#65471B",cx:"13.119",cy:"11.174",rx:"2.125",ry:"2.656"}),d("ellipse",{fill:"#65471B",cx:"24.375",cy:"12.236",rx:"2.125",ry:"2.656"}),d("path",{fill:"#F19020",d:"M17.276 35.149s1.265-.411 1.429-1.352c.173-.972-.624-1.167-.624-1.167s1.041-.208 1.172-1.376c.123-1.101-.861-1.363-.861-1.363s.97-.4 1.016-1.539c.038-.959-.995-1.428-.995-1.428s5.038-1.221 5.556-1.341c.516-.12 1.32-.615 1.069-1.694-.249-1.08-1.204-1.118-1.697-1.003-.494.115-6.744 1.566-8.9 2.068l-1.439.334c-.54.127-.785-.11-.404-.512.508-.536.833-1.129.946-2.113.119-1.035-.232-2.313-.433-2.809-.374-.921-1.005-1.649-1.734-1.899-1.137-.39-1.945.321-1.542 1.561.604 1.854.208 3.375-.833 4.293-2.449 2.157-3.588 3.695-2.83 6.973.828 3.575 4.377 5.876 7.952 5.048l3.152-.681z"}),d("path",{fill:"#65471B",d:"M9.296 6.351c-.164-.088-.303-.224-.391-.399-.216-.428-.04-.927.393-1.112 4.266-1.831 7.699-.043 7.843.034.433.231.608.747.391 1.154-.216.405-.74.546-1.173.318-.123-.063-2.832-1.432-6.278.047-.257.109-.547.085-.785-.042zm12.135 3.75c-.156-.098-.286-.243-.362-.424-.187-.442.023-.927.468-1.084 4.381-1.536 7.685.48 7.823.567.415.26.555.787.312 1.178-.242.39-.776.495-1.191.238-.12-.072-2.727-1.621-6.267-.379-.266.091-.553.046-.783-.096z"})],-1))})()}function zt(){return(()=>{const e=Z("de99a3ee9dd8b632");return e[0]||(e[0]=d("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 36 36"},[d("ellipse",{fill:"#292F33",cx:"18",cy:"26",rx:"18",ry:"10"}),d("ellipse",{fill:"#66757F",cx:"18",cy:"24",rx:"18",ry:"10"}),d("path",{fill:"#E1E8ED",d:"M18 31C3.042 31 1 16 1 12h34c0 2-1.958 19-17 19z"}),d("path",{fill:"#77B255",d:"M35 12.056c0 5.216-7.611 9.444-17 9.444S1 17.271 1 12.056C1 6.84 8.611 3.611 18 3.611s17 3.229 17 8.445z"}),d("ellipse",{fill:"#A6D388",cx:"18",cy:"13",rx:"15",ry:"7"}),d("path",{d:"M21 17c-.256 0-.512-.098-.707-.293-2.337-2.337-2.376-4.885-.125-8.262.739-1.109.9-2.246.478-3.377-.461-1.236-1.438-1.996-1.731-2.077-.553 0-.958-.443-.958-.996 0-.552.491-.995 1.043-.995.997 0 2.395 1.153 3.183 2.625 1.034 1.933.91 4.039-.351 5.929-1.961 2.942-1.531 4.332-.125 5.738.391.391.391 1.023 0 1.414-.195.196-.451.294-.707.294zm-6-2c-.256 0-.512-.098-.707-.293-2.337-2.337-2.376-4.885-.125-8.262.727-1.091.893-2.083.494-2.947-.444-.961-1.431-1.469-1.684-1.499-.552 0-.989-.447-.989-1 0-.552.458-1 1.011-1 .997 0 2.585.974 3.36 2.423.481.899 1.052 2.761-.528 5.131-1.961 2.942-1.531 4.332-.125 5.738.391.391.391 1.023 0 1.414-.195.197-.451.295-.707.295z",fill:"#5C913B"})],-1))})()}function kt(){return(()=>{const e=Z("96a420cda0b7dc21");return e[0]||(e[0]=d("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 36 36"},[d("path",{fill:"#FFCC4D",d:"M36 18c0 9.941-8.059 18-18 18-9.94 0-18-8.059-18-18C0 8.06 8.06 0 18 0c9.941 0 18 8.06 18 18"}),d("ellipse",{fill:"#664500",cx:"18",cy:"27",rx:"5",ry:"6"}),d("path",{fill:"#664500",d:"M5.999 11c-.208 0-.419-.065-.599-.2-.442-.331-.531-.958-.2-1.4C8.462 5.05 12.816 5 13 5c.552 0 1 .448 1 1 0 .551-.445.998-.996 1-.155.002-3.568.086-6.204 3.6-.196.262-.497.4-.801.4zm24.002 0c-.305 0-.604-.138-.801-.4-2.64-3.521-6.061-3.598-6.206-3.6-.55-.006-.994-.456-.991-1.005C22.006 5.444 22.45 5 23 5c.184 0 4.537.05 7.8 4.4.332.442.242 1.069-.2 1.4-.18.135-.39.2-.599.2zm-16.087 4.5l1.793-1.793c.391-.391.391-1.023 0-1.414s-1.023-.391-1.414 0L12.5 14.086l-1.793-1.793c-.391-.391-1.023-.391-1.414 0s-.391 1.023 0 1.414l1.793 1.793-1.793 1.793c-.391.391-.391 1.023 0 1.414.195.195.451.293.707.293s.512-.098.707-.293l1.793-1.793 1.793 1.793c.195.195.451.293.707.293s.512-.098.707-.293c.391-.391.391-1.023 0-1.414L13.914 15.5zm11 0l1.793-1.793c.391-.391.391-1.023 0-1.414s-1.023-.391-1.414 0L23.5 14.086l-1.793-1.793c-.391-.391-1.023-.391-1.414 0s-.391 1.023 0 1.414l1.793 1.793-1.793 1.793c-.391.391-.391 1.023 0 1.414.195.195.451.293.707.293s.512-.098.707-.293l1.793-1.793 1.793 1.793c.195.195.451.293.707.293s.512-.098.707-.293c.391-.391.391-1.023 0-1.414L24.914 15.5z"})],-1))})()}var _t=x("result",`
 color: var(--n-text-color);
 line-height: var(--n-line-height);
 font-size: var(--n-font-size);
 transition:
 color .3s var(--n-bezier);
`,[x("result-icon",`
 display: flex;
 justify-content: center;
 transition: color .3s var(--n-bezier);
 `,[ee("status-image",`
 font-size: var(--n-icon-size);
 width: 1em;
 height: 1em;
 `),x("base-icon",`
 color: var(--n-icon-color);
 font-size: var(--n-icon-size);
 `)]),x("result-content",{marginTop:"24px"}),x("result-footer",`
 margin-top: 24px;
 text-align: center;
 `),x("result-header",[ee("title",`
 margin-top: 16px;
 font-weight: var(--n-title-font-weight);
 transition: color .3s var(--n-bezier);
 text-align: center;
 color: var(--n-title-text-color);
 font-size: var(--n-title-font-size);
 `),ee("description",`
 margin-top: 4px;
 text-align: center;
 font-size: var(--n-font-size);
 `)])]);const Rt={403:wt,404:$t,418:zt,500:kt,info:()=>(l(),S(le)),success:()=>(l(),S(de)),warning:()=>(l(),S(ae)),error:()=>(l(),S(ce))},Pt={...A.props,size:String,status:{type:String,default:"info"},title:String,description:String};var Bt=M({name:"Result",props:Pt,slots:Object,setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:o,mergedComponentPropsRef:a}=J(e),i=w(()=>e.size||a?.value?.Result?.size||"medium"),r=A("Result","-result",_t,Qe,e,t),n=w(()=>{const{status:c}=e,g=i.value,{common:{cubicBezierEaseInOut:v},self:{textColor:u,lineHeight:m,titleTextColor:C,titleFontWeight:p,[O("iconColor",c)]:f,[O("fontSize",g)]:k,[O("titleFontSize",g)]:$,[O("iconSize",g)]:_}}=r.value;return{"--n-bezier":v,"--n-font-size":k,"--n-icon-size":_,"--n-line-height":m,"--n-text-color":u,"--n-title-font-size":$,"--n-title-font-weight":p,"--n-title-text-color":C,"--n-icon-color":f||""}}),s=o?ye("result",w(()=>{const{status:c}=e,g=i.value;let v="";return g&&(v+=g[0]),c&&(v+=c[0]),v}),n,e):void 0;return{mergedClsPrefix:t,cssVars:o?void 0:n,themeClass:s?.themeClass,onRender:s?.onRender}},render(){const{status:e,$slots:t,mergedClsPrefix:o,onRender:a}=this;return a?.(),l(),h("div",{class:y([`${o}-result`,this.themeClass]),style:R(this.cssVars)},[d("div",{class:y(`${o}-result-icon`)},[b(()=>t.icon?.()||(l(),S(ue,{clsPrefix:o},{default:()=>Rt[e]()},1032,["clsPrefix"])))],2),d("div",{class:y(`${o}-result-header`)},[this.title?(l(),h("div",{key:0,class:y(`${o}-result-header__title`)},[b(()=>this.title)],2)):b(()=>null),this.description?(l(),h("div",{key:2,class:y(`${o}-result-header__description`)},[b(()=>this.description)],2)):b(()=>null)],2),b(()=>t.default&&(l(),h("div",{class:y(`${o}-result-content`)},[b(()=>t.default())],2))),b(()=>t.footer&&(l(),h("div",{class:y(`${o}-result-footer`)},[b(()=>t.footer())],2)))],6)}});function Nt(e){const{heightSmall:t,heightMedium:o,heightLarge:a,borderRadius:i}=e;return{color:"#eee",colorEnd:"#ddd",borderRadius:i,heightSmall:t,heightMedium:o,heightLarge:a}}const Dt={common:ne,self:Nt};var Gt=T([x("skeleton",`
 height: 1em;
 width: 100%;
 transition:
 --n-color-start .3s var(--n-bezier),
 --n-color-end .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 animation: 2s skeleton-loading infinite cubic-bezier(0.36, 0, 0.64, 1);
 background-color: var(--n-color-start);
 `),T("@keyframes skeleton-loading",`
 0% {
 background: var(--n-color-start);
 }
 40% {
 background: var(--n-color-end);
 }
 80% {
 background: var(--n-color-start);
 }
 100% {
 background: var(--n-color-start);
 }
 `)]);const Mt={...A.props,text:Boolean,round:Boolean,circle:Boolean,height:[String,Number],width:[String,Number],size:String,repeat:{type:Number,default:1},animated:{type:Boolean,default:!0},sharp:{type:Boolean,default:!0}};var Y=M({name:"Skeleton",inheritAttrs:!1,props:Mt,setup(e){Ze();const{mergedClsPrefixRef:t,mergedComponentPropsRef:o}=J(e),a=w(()=>e.size||o?.value?.Skeleton?.size),i=A("Skeleton","-skeleton",Gt,Dt,e,t);return{mergedClsPrefix:t,style:w(()=>{const r=i.value,{common:{cubicBezierEaseInOut:n}}=r,s=r.self,{color:c,colorEnd:g,borderRadius:v}=s;let u;const{circle:m,sharp:C,round:p,width:f,height:k,text:$,animated:_}=e,N=a.value;N!==void 0&&(u=s[O("height",N)]);const D=m?f??k??u:f,V=(m?f??k:k)??u;return{display:$?"inline-block":"",verticalAlign:$?"-0.125em":"",borderRadius:m?"50%":p?"4096px":C?"":v,width:typeof D=="number"?W(D):D,height:typeof V=="number"?W(V):V,animation:_?"":"none","--n-bezier":n,"--n-color-start":c,"--n-color-end":g}})}},render(){const{repeat:e,style:t,mergedClsPrefix:o,$attrs:a}=this,i=ie("div",se({class:`${o}-skeleton`,style:t},a));return e>1?(l(),h(B,{key:1},[b(()=>Me(e,null).map(r=>[i,`
`]))],64)):i}});const Et={class:"error"},It=M({__name:"ErrorState",props:{description:{default:"请稍后重试"}},emits:["retry"],setup(e,{emit:t}){const o=t;return(a,i)=>(l(),h("div",Et,[P(z(Bt),{status:"error",title:"加载失败",description:e.description},{footer:G(()=>[P(z(Ee),{size:"small",onClick:i[0]||(i[0]=r=>o("retry"))},{default:G(()=>[...i[1]||(i[1]=[Ie("重试",-1)])]),_:1})]),_:1},8,["description"])]))}}),Wt=xe(It,[["__scopeId","data-v-359f1030"]]),Vt={class:"metric"},Lt={class:"metric"},Ot={class:"metric"},jt={key:1,class:"rows"},Ft=M({__name:"DashboardView",setup(e){const t=L(null),o=L(!1);let a=null;function i(){o.value=!1,a?.close();const s=new EventSource("/api/admin/dashboard");a=s,s.addEventListener("dashboard",c=>{t.value=JSON.parse(c.data)}),s.onerror=()=>{s.close(),a=null,t.value||(o.value=!0)}}i(),We(()=>a?.close());function r(s){return typeof s!="number"||Number.isNaN(s)?0:Math.min(100,Math.max(0,Number(s.toFixed(1))))}const n=[["项目版本","ProjectVersion"],["运行环境","Env"],["操作系统","GoOS"],["架构","GoArch"],["Go 版本","GoVersion"],["主机名","Host"],["CPU","CpuName"],["核心数","CpuCores"],["项目路径","ProjectPath"]];return(s,c)=>o.value?(l(),S(Wt,{key:0,onRetry:i})):(l(),h(B,{key:1},[P(z(rt),{cols:3,"x-gap":12,"y-gap":12,responsive:"screen","item-responsive":""},{default:G(()=>[P(z(te),{span:"3 m:1"},{default:G(()=>[P(z(X),{size:"small",title:"内存"},{default:G(()=>[t.value?(l(),h(B,{key:1},[P(z(oe),{type:"line",percentage:r(t.value.MemUsedPercent),height:6,"show-indicator":!1},null,8,["percentage"]),d("p",Vt,I(t.value.MemUsed)+" / "+I(t.value.MemTotal),1)],64)):(l(),S(z(Y),{key:0,text:"",repeat:3}))]),_:1})]),_:1}),P(z(te),{span:"3 m:1"},{default:G(()=>[P(z(X),{size:"small",title:"磁盘"},{default:G(()=>[t.value?(l(),h(B,{key:1},[P(z(oe),{type:"line",percentage:r(t.value.DiskUsedPercent),height:6,"show-indicator":!1},null,8,["percentage"]),d("p",Lt,I(t.value.DiskUsed)+" / "+I(t.value.DiskTotal),1)],64)):(l(),S(z(Y),{key:0,text:"",repeat:3}))]),_:1})]),_:1}),P(z(te),{span:"3 m:1"},{default:G(()=>[P(z(X),{size:"small",title:"CPU"},{default:G(()=>[t.value?(l(),h(B,{key:1},[P(z(oe),{type:"line",percentage:r(t.value.CpuUsedPercent),height:6,"show-indicator":!1},null,8,["percentage"]),d("p",Ot,I(r(t.value.CpuUsedPercent))+"% · "+I(t.value.CpuCores)+" 核",1)],64)):(l(),S(z(Y),{key:0,text:"",repeat:3}))]),_:1})]),_:1})]),_:1}),P(z(X),{size:"small",title:"系统信息",class:"system"},{default:G(()=>[t.value?(l(),h("dl",jt,[(l(),h(B,null,Ve(n,([g,v])=>(l(),h(B,{key:v},[d("dt",null,I(g),1),d("dd",null,I(t.value[v]),1)],64))),64))])):(l(),S(z(Y),{key:0,text:"",repeat:6}))]),_:1})],64))}}),qt=xe(Ft,[["__scopeId","data-v-0508084e"]]);export{qt as default};
