import{n as V,d as q,o as r,c as s,a as u,z as S,A as l,C as y,O as D,D as $,J as w,aJ as M,aK as T,aL as A,aM as X,aN as F,N as E,M as I,v as g,x as B,y as Y,E as j,G as J,I as K,ac as G}from"./index-BGprsyhQ.js";import{f as N}from"./_plugin-vue_export-helper-CtyYQGco.js";function Z(e){const{infoColor:f,successColor:h,warningColor:m,errorColor:c,textColor2:t,progressRailColor:i,fontSize:d,fontWeight:p}=e;return{fontSize:d,fontSizeCircle:"28px",fontWeightCircle:p,railColor:i,railHeight:"8px",iconSizeCircle:"36px",iconSizeLine:"18px",iconColor:f,iconColorInfo:f,iconColorSuccess:h,iconColorWarning:m,iconColorError:c,textColorCircle:t,textColorLineInner:"rgb(255, 255, 255)",textColorLineOuter:t,fillColor:f,fillColorInfo:f,fillColorSuccess:h,fillColorWarning:m,fillColorError:c,lineBgProcessing:"linear-gradient(90deg, rgba(255, 255, 255, .3) 0%, rgba(255, 255, 255, .5) 100%)"}}const Q={name:"Progress",common:V,self:Z},U=["id"],ee=["stop-color"],re=["stop-color"],te=["viewBox"],oe=["d","stroke-width"],ie=["d","stroke-width"],le={success:(r(),$(X)),error:(r(),$(A)),warning:(r(),$(T)),info:(r(),$(M))};var se=q({name:"ProgressCircle",props:{clsPrefix:{type:String,required:!0},status:{type:String,required:!0},strokeWidth:{type:Number,required:!0},fillColor:[String,Object],railColor:String,railStyle:[String,Object],percentage:{type:Number,default:0},offsetDegree:{type:Number,default:0},showIndicator:{type:Boolean,required:!0},indicatorTextColor:String,unit:String,viewBoxWidth:{type:Number,required:!0},gapDegree:{type:Number,required:!0},gapOffsetDegree:{type:Number,default:0}},setup(e,{slots:f}){const h=w(()=>{const t="gradient",{fillColor:i}=e;return typeof i=="object"?`${t}-${F(JSON.stringify(i))}`:t});function m(t,i,d,p){const{gapDegree:v,viewBoxWidth:b,strokeWidth:x}=e,n=50,C=0,a=n,o=0,_=100,P=50+x/2,k=`M ${P},${P} m ${C},${a}
      a ${n},${n} 0 1 1 ${o},-100
      a ${n},${n} 0 1 1 0,${_}`,z=Math.PI*2*n;return{pathString:k,pathStyle:{stroke:p==="rail"?d:typeof e.fillColor=="object"?`url(#${h.value})`:d,strokeDasharray:`${Math.min(t,100)/100*(z-v)}px ${b*8}px`,strokeDashoffset:`-${v/2}px`,transformOrigin:i?"center":void 0,transform:i?`rotate(${i}deg)`:void 0}}}const c=()=>{const t=typeof e.fillColor=="object",i=t?e.fillColor.stops[0]:"",d=t?e.fillColor.stops[1]:"";return t&&(r(),s("defs",null,[u("linearGradient",{id:h.value,x1:"0%",y1:"100%",x2:"100%",y2:"0%"},[u("stop",{offset:"0%","stop-color":i},null,8,ee),u("stop",{offset:"100%","stop-color":d},null,8,re)],8,U)]))};return()=>{const{fillColor:t,railColor:i,strokeWidth:d,offsetDegree:p,status:v,percentage:b,showIndicator:x,indicatorTextColor:n,unit:C,gapOffsetDegree:a,clsPrefix:o}=e,{pathString:_,pathStyle:P}=m(100,0,i,"rail"),{pathString:k,pathStyle:z}=m(b,p,t,"fill"),R=100+d;return r(),s("div",{class:l(`${o}-progress-content`),role:"none"},[u("div",{class:l(`${o}-progress-graph`),"aria-hidden":!0},[u("div",{class:l(`${o}-progress-graph-circle`),style:S({transform:a?`rotate(${a}deg)`:void 0})},[(r(),s("svg",{viewBox:`0 0 ${R} ${R}`},[y(()=>c()),u("g",null,[u("path",{class:l(`${o}-progress-graph-circle-rail`),d:_,"stroke-width":d,"stroke-linecap":"round",fill:"none",style:S(P)},null,14,oe)]),u("g",null,[u("path",{class:l([`${o}-progress-graph-circle-fill`,b===0&&`${o}-progress-graph-circle-fill--empty`]),d:k,"stroke-width":d,"stroke-linecap":"round",fill:"none",style:S(z)},null,14,ie)])],8,te))],6)],2),x?(r(),s("div",{key:0},[f.default?(r(),s("div",{key:0,class:l(`${o}-progress-custom-content`),role:"none"},[y(()=>f.default())],2)):(r(),s(D,{key:1},[v!=="default"?(r(),s("div",{key:0,class:l(`${o}-progress-icon`),"aria-hidden":!0},[(r(),$(E,{clsPrefix:o},{default:()=>le[v]},1032,["clsPrefix"]))],2)):(r(),s("div",{key:1,class:l(`${o}-progress-text`),style:S({color:n}),role:"none"},[u("span",{class:l(`${o}-progress-text__percentage`)},[y(()=>b)],2),u("span",{class:l(`${o}-progress-text__unit`)},[y(()=>C)],2)],6))],64))])):y(()=>null)],2)}}});const ne={success:(r(),$(X)),error:(r(),$(A)),warning:(r(),$(T)),info:(r(),$(M))};var ae=q({name:"ProgressLine",props:{clsPrefix:{type:String,required:!0},percentage:{type:Number,default:0},railColor:String,railStyle:[String,Object],fillColor:[String,Object],status:{type:String,required:!0},indicatorPlacement:{type:String,required:!0},indicatorTextColor:String,unit:{type:String,default:"%"},processing:{type:Boolean,required:!0},showIndicator:{type:Boolean,required:!0},height:[String,Number],railBorderRadius:[String,Number],fillBorderRadius:[String,Number]},setup(e,{slots:f}){const h=w(()=>N(e.height)),m=w(()=>typeof e.fillColor=="object"?`linear-gradient(to right, ${e.fillColor?.stops[0]} , ${e.fillColor?.stops[1]})`:e.fillColor),c=w(()=>e.railBorderRadius!==void 0?N(e.railBorderRadius):e.height!==void 0?N(e.height,{c:.5}):""),t=w(()=>e.fillBorderRadius!==void 0?N(e.fillBorderRadius):e.railBorderRadius!==void 0?N(e.railBorderRadius):e.height!==void 0?N(e.height,{c:.5}):"");return()=>{const{indicatorPlacement:i,railColor:d,railStyle:p,percentage:v,unit:b,indicatorTextColor:x,status:n,showIndicator:C,processing:a,clsPrefix:o}=e;return r(),s("div",{class:l(`${o}-progress-content`),role:"none"},[u("div",{class:l(`${o}-progress-graph`),"aria-hidden":!0},[u("div",{class:l([`${o}-progress-graph-line`,{[`${o}-progress-graph-line--indicator-${i}`]:!0}])},[u("div",{class:l(`${o}-progress-graph-line-rail`),style:S([{backgroundColor:d,height:h.value,borderRadius:c.value},p])},[u("div",{class:l([`${o}-progress-graph-line-fill`,a&&`${o}-progress-graph-line-fill--processing`]),style:S({maxWidth:`${e.percentage}%`,background:m.value,height:h.value,lineHeight:h.value,borderRadius:t.value})},[i==="inside"?(r(),s("div",{key:0,class:l(`${o}-progress-graph-line-indicator`),style:S({color:x})},[f.default?(r(),s(D,{key:0},[y(()=>f.default())],64)):(r(),s(D,{key:1},[y(()=>`${v}${b}`)],64))],6)):y(()=>null)],6)],6)],2)],2),C&&i==="outside"?(r(),s("div",{key:0},[f.default?(r(),s("div",{key:0,class:l(`${o}-progress-custom-content`),style:S({color:x}),role:"none"},[y(()=>f.default())],6)):(r(),s(D,{key:1},[n==="default"?(r(),s("div",{key:0,role:"none",class:l(`${o}-progress-icon ${o}-progress-icon--as-text`),style:S({color:x})},[y(()=>v),y(()=>b)],6)):(r(),s("div",{key:1,class:l(`${o}-progress-icon`),"aria-hidden":!0},[(r(),$(E,{clsPrefix:o},{default:()=>ne[n]},1032,["clsPrefix"]))],2))],64))])):y(()=>null)],2)}}});const ce=["id"],de=["stop-color"],ge=["stop-color"],ue=["d","stroke-width"],fe=["d","stroke-width"],pe=["viewBox"];function L(e,f,h=100){return`m ${h/2} ${h/2-e} a ${e} ${e} 0 1 1 0 ${2*e} a ${e} ${e} 0 1 1 0 -${2*e}`}var he=q({name:"ProgressMultipleCircle",props:{clsPrefix:{type:String,required:!0},viewBoxWidth:{type:Number,required:!0},percentage:{type:Array,default:[0]},strokeWidth:{type:Number,required:!0},circleGap:{type:Number,required:!0},showIndicator:{type:Boolean,required:!0},fillColor:{type:Array,default:()=>[]},railColor:{type:Array,default:()=>[]},railStyle:{type:Array,default:()=>[]}},setup(e,{slots:f}){const h=w(()=>e.percentage.map((c,t)=>`${Math.PI*c/100*(e.viewBoxWidth/2-e.strokeWidth/2*(1+2*t)-e.circleGap*t)*2}, ${e.viewBoxWidth*8}`)),m=(c,t)=>{const i=e.fillColor[t],d=typeof i=="object"?i.stops[0]:"",p=typeof i=="object"?i.stops[1]:"";return typeof e.fillColor[t]=="object"&&(r(),s("linearGradient",{id:`gradient-${t}`,x1:"100%",y1:"0%",x2:"0%",y2:"100%"},[u("stop",{offset:"0%","stop-color":d},null,8,de),u("stop",{offset:"100%","stop-color":p},null,8,ge)],8,ce))};return()=>{const{viewBoxWidth:c,strokeWidth:t,circleGap:i,showIndicator:d,fillColor:p,railColor:v,railStyle:b,percentage:x,clsPrefix:n}=e;return r(),s("div",{class:l(`${n}-progress-content`),role:"none"},[u("div",{class:l(`${n}-progress-graph`),"aria-hidden":!0},[u("div",{class:l(`${n}-progress-graph-circle`)},[(r(),s("svg",{viewBox:`0 0 ${c} ${c}`},[u("defs",null,[y(()=>x.map((C,a)=>m(C,a)))]),y(()=>x.map((C,a)=>(r(),s("g",{key:a},[u("path",{class:l(`${n}-progress-graph-circle-rail`),d:L(c/2-t/2*(1+2*a)-i*a,t,c),"stroke-width":t,"stroke-linecap":"round",fill:"none",style:S([{strokeDashoffset:0,stroke:v[a]},b[a]])},null,14,ue),u("path",{class:l([`${n}-progress-graph-circle-fill`,C===0&&`${n}-progress-graph-circle-fill--empty`]),d:L(c/2-t/2*(1+2*a)-i*a,t,c),"stroke-width":t,"stroke-linecap":"round",fill:"none",style:S({strokeDasharray:h.value[a],strokeDashoffset:0,stroke:typeof p[a]=="object"?`url(#gradient-${a})`:p[a]})},null,14,fe)]))))],8,pe))],2)],2),d&&f.default?(r(),s("div",{key:0},[u("div",{class:l(`${n}-progress-text`)},[y(()=>f.default())],2)])):y(()=>null)],2)}}}),ye=I([g("progress",{display:"inline-block"},[g("progress-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 `),B("line",`
 width: 100%;
 display: block;
 `,[g("progress-content",`
 display: flex;
 align-items: center;
 `,[g("progress-graph",{flex:1})]),g("progress-custom-content",{marginLeft:"14px"}),g("progress-icon",`
 width: 30px;
 padding-left: 14px;
 height: var(--n-icon-size-line);
 line-height: var(--n-icon-size-line);
 font-size: var(--n-icon-size-line);
 `,[B("as-text",`
 color: var(--n-text-color-line-outer);
 text-align: center;
 width: 40px;
 font-size: var(--n-font-size);
 padding-left: 4px;
 transition: color .3s var(--n-bezier);
 `)])]),B("circle, dashboard",{width:"120px"},[g("progress-custom-content",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 display: flex;
 align-items: center;
 justify-content: center;
 `),g("progress-text",`
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
 `),g("progress-icon",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 display: flex;
 align-items: center;
 color: var(--n-icon-color);
 font-size: var(--n-icon-size-circle);
 `)]),B("multiple-circle",`
 width: 200px;
 color: inherit;
 `,[g("progress-text",`
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
 `)]),g("progress-content",{position:"relative"}),g("progress-graph",{position:"relative"},[g("progress-graph-circle",[I("svg",{verticalAlign:"bottom"}),g("progress-graph-circle-fill",`
 stroke: var(--n-fill-color);
 transition:
 opacity .3s var(--n-bezier),
 stroke .3s var(--n-bezier),
 stroke-dasharray .3s var(--n-bezier);
 `,[B("empty",{opacity:0})]),g("progress-graph-circle-rail",`
 transition: stroke .3s var(--n-bezier);
 overflow: hidden;
 stroke: var(--n-rail-color);
 `)]),g("progress-graph-line",[B("indicator-inside",[g("progress-graph-line-rail",`
 height: 16px;
 line-height: 16px;
 border-radius: 10px;
 `,[g("progress-graph-line-fill",`
 height: inherit;
 border-radius: 10px;
 `),g("progress-graph-line-indicator",`
 background: #0000;
 white-space: nowrap;
 text-align: right;
 margin-left: 14px;
 margin-right: 14px;
 height: inherit;
 font-size: 12px;
 color: var(--n-text-color-line-inner);
 transition: color .3s var(--n-bezier);
 `)])]),B("indicator-inside-label",`
 height: 16px;
 display: flex;
 align-items: center;
 `,[g("progress-graph-line-rail",`
 flex: 1;
 transition: background-color .3s var(--n-bezier);
 `),g("progress-graph-line-indicator",`
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
 `)]),g("progress-graph-line-rail",`
 position: relative;
 overflow: hidden;
 height: var(--n-rail-height);
 border-radius: 5px;
 background-color: var(--n-rail-color);
 transition: background-color .3s var(--n-bezier);
 `,[g("progress-graph-line-fill",`
 background: var(--n-fill-color);
 position: relative;
 border-radius: 5px;
 height: inherit;
 width: 100%;
 max-width: 0%;
 transition:
 background-color .3s var(--n-bezier),
 max-width .2s var(--n-bezier);
 `,[B("processing",[I("&::after",`
 content: "";
 background-image: var(--n-line-bg-processing);
 animation: progress-processing-animation 2s var(--n-bezier) infinite;
 `)])])])])])]),I("@keyframes progress-processing-animation",`
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
 `)]);const me=["aria-valuenow","role"],ve={...Y.props,processing:Boolean,type:{type:String,default:"line"},gapDegree:Number,gapOffsetDegree:Number,status:{type:String,default:"default"},railColor:[String,Array],railStyle:[String,Array],color:[String,Array,Object],viewBoxWidth:{type:Number,default:100},strokeWidth:{type:Number,default:7},percentage:[Number,Array],unit:{type:String,default:"%"},showIndicator:{type:Boolean,default:!0},indicatorPosition:{type:String,default:"outside"},indicatorPlacement:{type:String,default:"outside"},indicatorTextColor:String,circleGap:{type:Number,default:1},height:Number,borderRadius:[String,Number],fillBorderRadius:[String,Number],offsetDegree:Number};var Ce=q({name:"Progress",props:ve,setup(e){const f=w(()=>e.indicatorPlacement||e.indicatorPosition),h=w(()=>{if(e.gapDegree||e.gapDegree===0)return e.gapDegree;if(e.type==="dashboard")return 75}),{mergedClsPrefixRef:m,inlineThemeDisabled:c}=J(e),t=Y("Progress","-progress",ye,Q,e,m),i=w(()=>{const{status:p}=e,{common:{cubicBezierEaseInOut:v},self:{fontSize:b,fontSizeCircle:x,railColor:n,railHeight:C,iconSizeCircle:a,iconSizeLine:o,textColorCircle:_,textColorLineInner:P,textColorLineOuter:k,lineBgProcessing:z,fontWeightCircle:R,[G("iconColor",p)]:O,[G("fillColor",p)]:W}}=t.value;return{"--n-bezier":v,"--n-fill-color":W,"--n-font-size":b,"--n-font-size-circle":x,"--n-font-weight-circle":R,"--n-icon-color":O,"--n-icon-size-circle":a,"--n-icon-size-line":o,"--n-line-bg-processing":z,"--n-rail-color":n,"--n-rail-height":C,"--n-text-color-circle":_,"--n-text-color-line-inner":P,"--n-text-color-line-outer":k}}),d=c?K("progress",w(()=>e.status[0]),i,e):void 0;return{mergedClsPrefix:m,mergedIndicatorPlacement:f,gapDeg:h,cssVars:c?void 0:i,themeClass:d?.themeClass,onRender:d?.onRender}},render(){const{type:e,cssVars:f,indicatorTextColor:h,showIndicator:m,status:c,railColor:t,railStyle:i,color:d,percentage:p,viewBoxWidth:v,strokeWidth:b,mergedIndicatorPlacement:x,unit:n,borderRadius:C,fillBorderRadius:a,height:o,processing:_,circleGap:P,mergedClsPrefix:k,gapDeg:z,gapOffsetDegree:R,themeClass:O,$slots:W,onRender:H}=this;return H?.(),r(),s("div",{class:l([O,`${k}-progress`,`${k}-progress--${e}`,`${k}-progress--${c}`]),style:S(f),"aria-valuemax":100,"aria-valuemin":0,"aria-valuenow":p,role:e==="circle"||e==="line"||e==="dashboard"?"progressbar":"none"},[e==="circle"||e==="dashboard"?(r(),$(se,{key:0,clsPrefix:k,status:c,showIndicator:m,indicatorTextColor:h,railColor:t,fillColor:d,railStyle:i,offsetDegree:this.offsetDegree,percentage:p,viewBoxWidth:v,strokeWidth:b,gapDegree:z===void 0?e==="dashboard"?75:0:z,gapOffsetDegree:R,unit:n},j(W),1032,["clsPrefix","status","showIndicator","indicatorTextColor","railColor","fillColor","railStyle","offsetDegree","percentage","viewBoxWidth","strokeWidth","gapDegree","gapOffsetDegree","unit"])):(r(),s(D,{key:1},[e==="line"?(r(),$(ae,{key:0,clsPrefix:k,status:c,showIndicator:m,indicatorTextColor:h,railColor:t,fillColor:d,railStyle:i,percentage:p,processing:_,indicatorPlacement:x,unit:n,fillBorderRadius:a,railBorderRadius:C,height:o},j(W),1032,["clsPrefix","status","showIndicator","indicatorTextColor","railColor","fillColor","railStyle","percentage","processing","indicatorPlacement","unit","fillBorderRadius","railBorderRadius","height"])):(r(),s(D,{key:1},[e==="multiple-circle"?(r(),$(he,{key:0,clsPrefix:k,strokeWidth:b,railColor:t,fillColor:d,railStyle:i,viewBoxWidth:v,percentage:p,showIndicator:m,circleGap:P},j(W),1032,["clsPrefix","strokeWidth","railColor","fillColor","railStyle","viewBoxWidth","percentage","showIndicator","circleGap"])):y(()=>null)],64))],64))],14,me)}});export{Ce as P,Q as p};
