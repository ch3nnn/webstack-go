import{m as Le,n as Fe,p as me,s as uo,q as Pe,t as ie,v as d,x as S,d as $,y as Y,o as a,c as g,z as U,A as T,C as R,D as B,E as Ke,F as X,S as Ve,G as ue,H as je,I as ve,J as x,k as F,K as ne,L as s,M as z,a as y,N as De,O as D,P as W,Q as V,R as de,T as le,U as vo,V as we,W as te,X as ze,Y as xe,Z as ho,_ as re,$ as pe,a0 as mo,a1 as Te,a2 as po,a3 as fo,u as go,w as Z,f as K,b as J,a4 as xo,a5 as Co,a6 as Ne,a7 as bo,a8 as yo,a9 as wo,h as zo,l as Io,g as ko}from"./index-Df1q10-G.js";import{d as So,t as Ao,C as Ho,T as Ro,D as Ue,V as _o,c as fe,N as Po}from"./Dropdown-CnvFPVek.js";import{a as Ce,u as To}from"./use-message-fTzDVPVq.js";import{f as ge,_ as No}from"./_plugin-vue_export-helper-CtyYQGco.js";import{u as Bo}from"./use-compitable-CjWuX6-V.js";import"./next-frame-once-C5Ksf8W7.js";function $o(e,t,o,n){return{itemColorHoverInverted:"#0000",itemColorActiveInverted:t,itemColorActiveHoverInverted:t,itemColorActiveCollapsedInverted:t,itemTextColorInverted:e,itemTextColorHoverInverted:o,itemTextColorChildActiveInverted:o,itemTextColorChildActiveHoverInverted:o,itemTextColorActiveInverted:o,itemTextColorActiveHoverInverted:o,itemTextColorHorizontalInverted:e,itemTextColorHoverHorizontalInverted:o,itemTextColorChildActiveHorizontalInverted:o,itemTextColorChildActiveHoverHorizontalInverted:o,itemTextColorActiveHorizontalInverted:o,itemTextColorActiveHoverHorizontalInverted:o,itemIconColorInverted:e,itemIconColorHoverInverted:o,itemIconColorActiveInverted:o,itemIconColorActiveHoverInverted:o,itemIconColorChildActiveInverted:o,itemIconColorChildActiveHoverInverted:o,itemIconColorCollapsedInverted:e,itemIconColorHorizontalInverted:e,itemIconColorHoverHorizontalInverted:o,itemIconColorActiveHorizontalInverted:o,itemIconColorActiveHoverHorizontalInverted:o,itemIconColorChildActiveHorizontalInverted:o,itemIconColorChildActiveHoverHorizontalInverted:o,arrowColorInverted:e,arrowColorHoverInverted:o,arrowColorActiveInverted:o,arrowColorActiveHoverInverted:o,arrowColorChildActiveInverted:o,arrowColorChildActiveHoverInverted:o,groupTextColorInverted:n}}function Mo(e){const{borderRadius:t,textColor3:o,primaryColor:n,textColor2:c,textColor1:i,fontSize:u,dividerColor:h,hoverColor:v,primaryColorHover:m}=e;return{borderRadius:t,color:"#0000",groupTextColor:o,itemColorHover:v,itemColorActive:me(n,{alpha:.1}),itemColorActiveHover:me(n,{alpha:.1}),itemColorActiveCollapsed:me(n,{alpha:.1}),itemTextColor:c,itemTextColorHover:c,itemTextColorActive:n,itemTextColorActiveHover:n,itemTextColorChildActive:n,itemTextColorChildActiveHover:n,itemTextColorHorizontal:c,itemTextColorHoverHorizontal:m,itemTextColorActiveHorizontal:n,itemTextColorActiveHoverHorizontal:n,itemTextColorChildActiveHorizontal:n,itemTextColorChildActiveHoverHorizontal:n,itemIconColor:i,itemIconColorHover:i,itemIconColorActive:n,itemIconColorActiveHover:n,itemIconColorChildActive:n,itemIconColorChildActiveHover:n,itemIconColorCollapsed:i,itemIconColorHorizontal:i,itemIconColorHoverHorizontal:m,itemIconColorActiveHorizontal:n,itemIconColorActiveHoverHorizontal:n,itemIconColorChildActiveHorizontal:n,itemIconColorChildActiveHoverHorizontal:n,itemHeight:"42px",arrowColor:c,arrowColorHover:c,arrowColorActive:n,arrowColorActiveHover:n,arrowColorChildActive:n,arrowColorChildActiveHover:n,colorInverted:"#0000",borderColorHorizontal:"#0000",fontSize:u,dividerColor:h,...$o("#BBB",n,"#FFF","#AAA")}}const Oo=Le({name:"Menu",common:Fe,peers:{Tooltip:Ao,Dropdown:So},self:Mo});function Eo(e){const{baseColor:t,textColor2:o,bodyColor:n,cardColor:c,dividerColor:i,actionColor:u,scrollbarColor:h,scrollbarColorHover:v,invertedColor:m}=e;return{textColor:o,textColorInverted:"#FFF",color:n,colorEmbedded:u,headerColor:c,headerColorInverted:m,footerColor:u,footerColorInverted:m,headerBorderColor:i,headerBorderColorInverted:m,footerBorderColor:i,footerBorderColorInverted:m,siderBorderColor:i,siderBorderColorInverted:m,siderColor:c,siderColorInverted:m,siderToggleButtonBorder:`1px solid ${i}`,siderToggleButtonColor:t,siderToggleButtonIconColor:o,siderToggleButtonIconColorInverted:o,siderToggleBarColor:Pe(n,h),siderToggleBarColorHover:Pe(n,v),__invertScrollbar:"true"}}const Ie=Le({name:"Layout",common:Fe,peers:{Scrollbar:uo},self:Eo}),Ge=ie("n-layout-sider"),ke={type:String,default:"static"};var Lo=d("layout",`
 color: var(--n-text-color);
 background-color: var(--n-color);
 box-sizing: border-box;
 position: relative;
 z-index: auto;
 flex: auto;
 overflow: hidden;
 transition:
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
`,[d("layout-scroll-container",`
 overflow-x: hidden;
 box-sizing: border-box;
 height: 100%;
 `),S("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]);const Fo={embedded:Boolean,position:ke,nativeScrollbar:{type:Boolean,default:!0},scrollbarProps:Object,onScroll:Function,contentClass:String,contentStyle:{type:[String,Object],default:""},hasSider:Boolean,siderPlacement:{type:String,default:"left"}},qe=ie("n-layout");function We(e){return $({name:e?"LayoutContent":"Layout",props:{...Y.props,...Fo},setup(t){const o=F(null),n=F(null),{mergedClsPrefixRef:c,inlineThemeDisabled:i}=ue(t),u=Y("Layout","-layout",Lo,Ie,t,c);function h(b,P){if(t.nativeScrollbar){const{value:M}=o;M&&(P===void 0?M.scrollTo(b):M.scrollTo(b,P))}else{const{value:M}=n;M&&M.scrollTo(b,P)}}ne(qe,t);let v=0,m=0;const _=b=>{const P=b.target;v=P.scrollLeft,m=P.scrollTop,t.onScroll?.(b)};je(()=>{if(t.nativeScrollbar){const b=o.value;b&&(b.scrollTop=m,b.scrollLeft=v)}});const A={display:"flex",flexWrap:"nowrap",width:"100%",flexDirection:"row"},I={scrollTo:h},w=x(()=>{const{common:{cubicBezierEaseInOut:b},self:P}=u.value;return{"--n-bezier":b,"--n-color":t.embedded?P.colorEmbedded:P.color,"--n-text-color":P.textColor}}),C=i?ve("layout",x(()=>t.embedded?"e":""),w,t):void 0;return{mergedClsPrefix:c,scrollableElRef:o,scrollbarInstRef:n,hasSiderStyle:A,mergedTheme:u,handleNativeElScroll:_,cssVars:i?void 0:w,themeClass:C?.themeClass,onRender:C?.onRender,...I}},render(){const{mergedClsPrefix:t,hasSider:o}=this;this.onRender?.();const n=o?this.hasSiderStyle:void 0,c=[this.themeClass,e&&`${t}-layout-content`,`${t}-layout`,`${t}-layout--${this.position}-positioned`];return a(),g("div",{class:T(c),style:U(this.cssVars)},[this.nativeScrollbar?(a(),g("div",{key:0,ref:"scrollableElRef",class:T([`${t}-layout-scroll-container`,this.contentClass]),style:U([this.contentStyle,n]),onScroll:this.handleNativeElScroll},[R(()=>this.$slots.default?.())],46,["onScroll"])):(a(),B(Ve,X({key:1},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,contentClass:this.contentClass,contentStyle:[this.contentStyle,n]}),Ke(this.$slots),1040,["onScroll","theme","themeOverrides","contentClass","contentStyle"]))],6)}})}var Be=We(!1),Ko=We(!0),Vo=d("layout-header",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 box-sizing: border-box;
 width: 100%;
 background-color: var(--n-color);
 color: var(--n-text-color);
`,[S("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 `),S("bordered",`
 border-bottom: solid 1px var(--n-border-color);
 `)]);const jo={position:ke,inverted:Boolean,bordered:Boolean};var Do=$({name:"LayoutHeader",props:{...Y.props,...jo},setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:o}=ue(e),n=Y("Layout","-layout-header",Vo,Ie,e,t),c=x(()=>{const{common:{cubicBezierEaseInOut:u},self:h}=n.value,v={"--n-bezier":u};return e.inverted?(v["--n-color"]=h.headerColorInverted,v["--n-text-color"]=h.textColorInverted,v["--n-border-color"]=h.headerBorderColorInverted):(v["--n-color"]=h.headerColor,v["--n-text-color"]=h.textColor,v["--n-border-color"]=h.headerBorderColor),v}),i=o?ve("layout-header",x(()=>e.inverted?"a":"b"),c,e):void 0;return{mergedClsPrefix:t,cssVars:o?void 0:c,themeClass:i?.themeClass,onRender:i?.onRender}},render(){const{mergedClsPrefix:e}=this;return this.onRender?.(),a(),g("div",{class:T([`${e}-layout-header`,this.themeClass,this.position&&`${e}-layout-header--${this.position}-positioned`,this.bordered&&`${e}-layout-header--bordered`]),style:U(this.cssVars)},[R(()=>this.$slots.default?.())],6)}}),Uo=d("layout-sider",`
 flex-shrink: 0;
 box-sizing: border-box;
 position: relative;
 z-index: 1;
 color: var(--n-text-color);
 transition:
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 min-width .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 transform .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 background-color: var(--n-color);
 display: flex;
 justify-content: flex-end;
`,[S("bordered",[s("border",`
 content: "";
 position: absolute;
 top: 0;
 bottom: 0;
 width: 1px;
 background-color: var(--n-border-color);
 transition: background-color .3s var(--n-bezier);
 `)]),s("left-placement",[S("bordered",[s("border",`
 right: 0;
 `)])]),S("right-placement",`
 justify-content: flex-start;
 `,[S("bordered",[s("border",`
 left: 0;
 `)]),S("collapsed",[d("layout-toggle-button",[d("base-icon",`
 transform: rotate(180deg);
 `)]),d("layout-toggle-bar",[z("&:hover",[s("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),s("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])])]),d("layout-toggle-button",`
 left: 0;
 transform: translateX(-50%) translateY(-50%);
 `,[d("base-icon",`
 transform: rotate(0);
 `)]),d("layout-toggle-bar",`
 left: -28px;
 transform: rotate(180deg);
 `,[z("&:hover",[s("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),s("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})])])]),S("collapsed",[d("layout-toggle-bar",[z("&:hover",[s("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),s("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])]),d("layout-toggle-button",[d("base-icon",`
 transform: rotate(0);
 `)])]),d("layout-toggle-button",`
 transition:
 color .3s var(--n-bezier),
 right .3s var(--n-bezier),
 left .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 cursor: pointer;
 width: 24px;
 height: 24px;
 position: absolute;
 top: 50%;
 right: 0;
 border-radius: 50%;
 display: flex;
 align-items: center;
 justify-content: center;
 font-size: 18px;
 color: var(--n-toggle-button-icon-color);
 border: var(--n-toggle-button-border);
 background-color: var(--n-toggle-button-color);
 box-shadow: 0 2px 4px 0px rgba(0, 0, 0, .06);
 transform: translateX(50%) translateY(-50%);
 z-index: 1;
 `,[d("base-icon",`
 transition: transform .3s var(--n-bezier);
 transform: rotate(180deg);
 `)]),d("layout-toggle-bar",`
 cursor: pointer;
 height: 72px;
 width: 32px;
 position: absolute;
 top: calc(50% - 36px);
 right: -28px;
 `,[s("top, bottom",`
 position: absolute;
 width: 4px;
 border-radius: 2px;
 height: 38px;
 left: 14px;
 transition: 
 background-color .3s var(--n-bezier),
 transform .3s var(--n-bezier);
 `),s("bottom",`
 position: absolute;
 top: 34px;
 `),z("&:hover",[s("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),s("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})]),s("top, bottom",{backgroundColor:"var(--n-toggle-bar-color)"}),z("&:hover",[s("top, bottom",{backgroundColor:"var(--n-toggle-bar-color-hover)"})])]),s("border",`
 position: absolute;
 top: 0;
 right: 0;
 bottom: 0;
 width: 1px;
 transition: background-color .3s var(--n-bezier);
 `),d("layout-sider-scroll-container",`
 flex-grow: 1;
 flex-shrink: 0;
 box-sizing: border-box;
 height: 100%;
 opacity: 0;
 transition: opacity .3s var(--n-bezier);
 max-width: 100%;
 `),S("show-content",[d("layout-sider-scroll-container",{opacity:1})]),S("absolute-positioned",`
 position: absolute;
 left: 0;
 top: 0;
 bottom: 0;
 `)]);const Go=["onClick"];var qo=$({props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){const{clsPrefix:e}=this;return a(),g("div",{onClick:this.onClick,class:T(`${e}-layout-toggle-bar`)},[y("div",{class:T(`${e}-layout-toggle-bar__top`)},null,2),y("div",{class:T(`${e}-layout-toggle-bar__bottom`)},null,2)],10,Go)}});const Wo=["onClick"];var Yo=$({name:"LayoutToggleButton",props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){const{clsPrefix:e}=this;return a(),g("div",{class:T(`${e}-layout-toggle-button`),onClick:this.onClick},[(a(),B(De,{clsPrefix:e},{default:()=>(a(),B(Ho))},1032,["clsPrefix"]))],10,Wo)}});const Xo=["onTransitionend"],Zo={position:ke,bordered:Boolean,collapsedWidth:{type:Number,default:48},width:{type:[Number,String],default:272},contentClass:String,contentStyle:{type:[String,Object],default:""},collapseMode:{type:String,default:"transform"},collapsed:{type:Boolean,default:void 0},defaultCollapsed:Boolean,showCollapsedContent:{type:Boolean,default:!0},showTrigger:{type:[Boolean,String],default:!1},nativeScrollbar:{type:Boolean,default:!0},inverted:Boolean,scrollbarProps:Object,triggerClass:String,triggerStyle:[String,Object],collapsedTriggerClass:String,collapsedTriggerStyle:[String,Object],"onUpdate:collapsed":[Function,Array],onUpdateCollapsed:[Function,Array],onAfterEnter:Function,onAfterLeave:Function,onExpand:[Function,Array],onCollapse:[Function,Array],onScroll:Function};var Jo=$({name:"LayoutSider",props:{...Y.props,...Zo},setup(e){const t=W(qe),o=F(null),n=F(null),c=F(e.defaultCollapsed),i=Ce(de(e,"collapsed"),c),u=x(()=>ge(i.value?e.collapsedWidth:e.width)),h=x(()=>e.collapseMode!=="transform"?{}:{minWidth:ge(e.width)}),v=x(()=>t?t.siderPlacement:"left");function m(N,H){if(e.nativeScrollbar){const{value:O}=o;O&&(H===void 0?O.scrollTo(N):O.scrollTo(N,H))}else{const{value:O}=n;O&&O.scrollTo(N,H)}}function _(){const{"onUpdate:collapsed":N,onUpdateCollapsed:H,onExpand:O,onCollapse:q}=e,{value:j}=i;H&&V(H,!j),N&&V(N,!j),c.value=!j,j?O&&V(O):q&&V(q)}let A=0,I=0;const w=N=>{const H=N.target;A=H.scrollLeft,I=H.scrollTop,e.onScroll?.(N)};je(()=>{if(e.nativeScrollbar){const N=o.value;N&&(N.scrollTop=I,N.scrollLeft=A)}}),ne(Ge,{collapsedRef:i,collapseModeRef:de(e,"collapseMode")});const{mergedClsPrefixRef:C,inlineThemeDisabled:b}=ue(e),P=Y("Layout","-layout-sider",Uo,Ie,e,C);function M(N){N.propertyName==="max-width"&&(i.value?e.onAfterLeave?.():e.onAfterEnter?.())}const ee={scrollTo:m},G=x(()=>{const{common:{cubicBezierEaseInOut:N},self:H}=P.value,{siderToggleButtonColor:O,siderToggleButtonBorder:q,siderToggleBarColor:j,siderToggleBarColorHover:he}=H,E={"--n-bezier":N,"--n-toggle-button-color":O,"--n-toggle-button-border":q,"--n-toggle-bar-color":j,"--n-toggle-bar-color-hover":he};return e.inverted?(E["--n-color"]=H.siderColorInverted,E["--n-text-color"]=H.textColorInverted,E["--n-border-color"]=H.siderBorderColorInverted,E["--n-toggle-button-icon-color"]=H.siderToggleButtonIconColorInverted,E.__invertScrollbar=H.__invertScrollbar):(E["--n-color"]=H.siderColor,E["--n-text-color"]=H.textColor,E["--n-border-color"]=H.siderBorderColor,E["--n-toggle-button-icon-color"]=H.siderToggleButtonIconColor),E}),oe=b?ve("layout-sider",x(()=>e.inverted?"a":"b"),G,e):void 0;return{scrollableElRef:o,scrollbarInstRef:n,mergedClsPrefix:C,mergedTheme:P,styleMaxWidth:u,mergedCollapsed:i,scrollContainerStyle:h,siderPlacement:v,handleNativeElScroll:w,handleTransitionend:M,handleTriggerClick:_,inlineThemeDisabled:b,cssVars:G,themeClass:oe?.themeClass,onRender:oe?.onRender,...ee}},render(){const{mergedClsPrefix:e,mergedCollapsed:t,showTrigger:o}=this;return this.onRender?.(),a(),g("aside",{class:T([`${e}-layout-sider`,this.themeClass,`${e}-layout-sider--${this.position}-positioned`,`${e}-layout-sider--${this.siderPlacement}-placement`,this.bordered&&`${e}-layout-sider--bordered`,t&&`${e}-layout-sider--collapsed`,(!t||this.showCollapsedContent)&&`${e}-layout-sider--show-content`]),onTransitionend:this.handleTransitionend,style:U([this.inlineThemeDisabled?void 0:this.cssVars,{maxWidth:this.styleMaxWidth,width:ge(this.width)}])},[this.nativeScrollbar?(a(),g("div",{key:1,class:T([`${e}-layout-sider-scroll-container`,this.contentClass]),onScroll:this.handleNativeElScroll,style:U([this.scrollContainerStyle,{overflow:"auto"},this.contentStyle]),ref:"scrollableElRef"},[R(()=>this.$slots.default?.())],46,["onScroll"])):(a(),B(Ve,X({key:0},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",style:this.scrollContainerStyle,contentStyle:this.contentStyle,contentClass:this.contentClass,theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,builtinThemeOverrides:this.inverted&&this.cssVars.__invertScrollbar==="true"?{colorHover:"rgba(255, 255, 255, .4)",color:"rgba(255, 255, 255, .3)"}:void 0}),Ke(this.$slots),1040,["onScroll","style","contentStyle","contentClass","theme","themeOverrides","builtinThemeOverrides"])),o?(a(),g(D,{key:2},[o==="bar"?(a(),B(qo,{key:0,clsPrefix:e,class:T(t?this.collapsedTriggerClass:this.triggerClass),style:U(t?this.collapsedTriggerStyle:this.triggerStyle),onClick:this.handleTriggerClick},null,8,["clsPrefix","class","style","onClick"])):(a(),B(Yo,{key:1,clsPrefix:e,class:T(t?this.collapsedTriggerClass:this.triggerClass),style:U(t?this.collapsedTriggerStyle:this.triggerStyle),onClick:this.handleTriggerClick},null,8,["clsPrefix","class","style","onClick"]))],64)):R(()=>null),this.bordered?(a(),g("div",{key:4,class:T(`${e}-layout-sider__border`)},null,2)):R(()=>null)],46,Xo)}});const ae=ie("n-menu"),Ye=ie("n-submenu"),Se=ie("n-menu-item-group"),$e=[z("&::before","background-color: var(--n-item-color-hover);"),s("arrow",`
 color: var(--n-arrow-color-hover);
 `),s("icon",`
 color: var(--n-item-icon-color-hover);
 `),d("menu-item-content-header",`
 color: var(--n-item-text-color-hover);
 `,[z("a",`
 color: var(--n-item-text-color-hover);
 `),s("extra",`
 color: var(--n-item-text-color-hover);
 `)])],Me=[s("icon",`
 color: var(--n-item-icon-color-hover-horizontal);
 `),d("menu-item-content-header",`
 color: var(--n-item-text-color-hover-horizontal);
 `,[z("a",`
 color: var(--n-item-text-color-hover-horizontal);
 `),s("extra",`
 color: var(--n-item-text-color-hover-horizontal);
 `)])];var Qo=z([d("menu",`
 background-color: var(--n-color);
 color: var(--n-item-text-color);
 overflow: hidden;
 transition: background-color .3s var(--n-bezier);
 box-sizing: border-box;
 font-size: var(--n-font-size);
 padding-bottom: 6px;
 `,[S("horizontal",`
 max-width: 100%;
 width: 100%;
 display: flex;
 overflow: hidden;
 padding-bottom: 0;
 `,[d("submenu","margin: 0;"),d("menu-item","margin: 0;"),d("menu-item-content",`
 padding: 0 20px;
 border-bottom: 2px solid #0000;
 `,[z("&::before","display: none;"),S("selected","border-bottom: 2px solid var(--n-border-color-horizontal)")]),d("menu-item-content",[S("selected",[s("icon","color: var(--n-item-icon-color-active-horizontal);"),d("menu-item-content-header",`
 color: var(--n-item-text-color-active-horizontal);
 `,[z("a","color: var(--n-item-text-color-active-horizontal);"),s("extra","color: var(--n-item-text-color-active-horizontal);")])]),S("child-active",`
 border-bottom: 2px solid var(--n-border-color-horizontal);
 `,[d("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-horizontal);
 `,[z("a",`
 color: var(--n-item-text-color-child-active-horizontal);
 `),s("extra",`
 color: var(--n-item-text-color-child-active-horizontal);
 `)]),s("icon",`
 color: var(--n-item-icon-color-child-active-horizontal);
 `)]),le("disabled",[le("selected, child-active",[z("&:focus-within",Me)]),S("selected",[Q(null,[s("icon","color: var(--n-item-icon-color-active-hover-horizontal);"),d("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover-horizontal);
 `,[z("a","color: var(--n-item-text-color-active-hover-horizontal);"),s("extra","color: var(--n-item-text-color-active-hover-horizontal);")])])]),S("child-active",[Q(null,[s("icon","color: var(--n-item-icon-color-child-active-hover-horizontal);"),d("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover-horizontal);
 `,[z("a","color: var(--n-item-text-color-child-active-hover-horizontal);"),s("extra","color: var(--n-item-text-color-child-active-hover-horizontal);")])])]),Q("border-bottom: 2px solid var(--n-border-color-horizontal);",Me)]),d("menu-item-content-header",[z("a","color: var(--n-item-text-color-horizontal);")])])]),le("responsive",[d("menu-item-content-header",`
 overflow: hidden;
 text-overflow: ellipsis;
 `)]),S("collapsed",[d("menu-item-content",[S("selected",[z("&::before",`
 background-color: var(--n-item-color-active-collapsed) !important;
 `)]),d("menu-item-content-header","opacity: 0;"),s("arrow","opacity: 0;"),s("icon","color: var(--n-item-icon-color-collapsed);")])]),d("menu-item",`
 height: var(--n-item-height);
 margin-top: 6px;
 position: relative;
 `),d("menu-item-content",`
 box-sizing: border-box;
 line-height: 1.75;
 height: 100%;
 display: grid;
 grid-template-areas: "icon content arrow";
 grid-template-columns: auto 1fr auto;
 align-items: center;
 cursor: pointer;
 position: relative;
 padding-right: 18px;
 transition:
 background-color .3s var(--n-bezier),
 padding-left .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[z("> *","z-index: 1;"),z("&::before",`
 z-index: auto;
 content: "";
 background-color: #0000;
 position: absolute;
 left: 8px;
 right: 8px;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border-radius: var(--n-border-radius);
 transition: background-color .3s var(--n-bezier);
 `),S("disabled",`
 opacity: .45;
 cursor: not-allowed;
 `),S("collapsed",[s("arrow","transform: rotate(0);")]),S("selected",[z("&::before","background-color: var(--n-item-color-active);"),s("arrow","color: var(--n-arrow-color-active);"),s("icon","color: var(--n-item-icon-color-active);"),d("menu-item-content-header",`
 color: var(--n-item-text-color-active);
 `,[z("a","color: var(--n-item-text-color-active);"),s("extra","color: var(--n-item-text-color-active);")])]),S("child-active",[d("menu-item-content-header",`
 color: var(--n-item-text-color-child-active);
 `,[z("a",`
 color: var(--n-item-text-color-child-active);
 `),s("extra",`
 color: var(--n-item-text-color-child-active);
 `)]),s("arrow",`
 color: var(--n-arrow-color-child-active);
 `),s("icon",`
 color: var(--n-item-icon-color-child-active);
 `)]),le("disabled",[le("selected, child-active",[z("&:focus-within",$e)]),S("selected",[Q(null,[s("arrow","color: var(--n-arrow-color-active-hover);"),s("icon","color: var(--n-item-icon-color-active-hover);"),d("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover);
 `,[z("a","color: var(--n-item-text-color-active-hover);"),s("extra","color: var(--n-item-text-color-active-hover);")])])]),S("child-active",[Q(null,[s("arrow","color: var(--n-arrow-color-child-active-hover);"),s("icon","color: var(--n-item-icon-color-child-active-hover);"),d("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover);
 `,[z("a","color: var(--n-item-text-color-child-active-hover);"),s("extra","color: var(--n-item-text-color-child-active-hover);")])])]),S("selected",[Q(null,[z("&::before","background-color: var(--n-item-color-active-hover);")])]),Q(null,$e)]),s("icon",`
 grid-area: icon;
 color: var(--n-item-icon-color);
 transition:
 color .3s var(--n-bezier),
 font-size .3s var(--n-bezier),
 margin-right .3s var(--n-bezier);
 box-sizing: content-box;
 display: inline-flex;
 align-items: center;
 justify-content: center;
 `),s("arrow",`
 grid-area: arrow;
 font-size: 16px;
 color: var(--n-arrow-color);
 transform: rotate(180deg);
 opacity: 1;
 transition:
 color .3s var(--n-bezier),
 transform 0.2s var(--n-bezier),
 opacity 0.2s var(--n-bezier);
 `),d("menu-item-content-header",`
 grid-area: content;
 transition:
 color .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 opacity: 1;
 white-space: nowrap;
 color: var(--n-item-text-color);
 `,[z("a",`
 outline: none;
 text-decoration: none;
 transition: color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 `,[z("&::before",`
 content: "";
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),s("extra",`
 font-size: .93em;
 color: var(--n-group-text-color);
 transition: color .3s var(--n-bezier);
 `)])]),d("submenu",`
 cursor: pointer;
 position: relative;
 margin-top: 6px;
 `,[d("menu-item-content",`
 height: var(--n-item-height);
 `),d("submenu-children",`
 overflow: hidden;
 padding: 0;
 `,[vo({duration:".2s"})])]),d("menu-item-group",[d("menu-item-group-title",`
 margin-top: 6px;
 color: var(--n-group-text-color);
 cursor: default;
 font-size: .93em;
 height: 36px;
 display: flex;
 align-items: center;
 transition:
 padding-left .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `)])]),d("menu-tooltip",[z("a",`
 color: inherit;
 text-decoration: none;
 `)]),d("menu-divider",`
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-divider-color);
 height: 1px;
 margin: 6px 18px;
 `)]);function Q(e,t){return[S("hover",e,t),z("&:hover",e,t)]}var et=$({name:"MenuDivider",setup(){const{mergedClsPrefixRef:e,isHorizontalRef:t}=W(ae);return()=>t.value?null:(a(),g("div",{key:1,class:T(`${e.value}-menu-divider`)},null,2))}}),ot=$({name:"ChevronDownFilled",render(){return(()=>{const e=we("f3af82a2aab086a5");return e[0]||(e[0]=y("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},[y("path",{d:"M3.20041 5.73966C3.48226 5.43613 3.95681 5.41856 4.26034 5.70041L8 9.22652L11.7397 5.70041C12.0432 5.41856 12.5177 5.43613 12.7996 5.73966C13.0815 6.0432 13.0639 6.51775 12.7603 6.7996L8.51034 10.7996C8.22258 11.0668 7.77743 11.0668 7.48967 10.7996L3.23966 6.7996C2.93613 6.51775 2.91856 6.0432 3.20041 5.73966Z",fill:"currentColor"})],-1))})()}});const tt=["onClick"];var Xe=$({name:"MenuOptionContent",props:{collapsed:Boolean,disabled:Boolean,title:[String,Function],icon:Function,extra:[String,Function],showArrow:Boolean,childActive:Boolean,hover:Boolean,paddingLeft:Number,selected:Boolean,maxIconSize:{type:Number,required:!0},activeIconSize:{type:Number,required:!0},iconMarginRight:{type:Number,required:!0},clsPrefix:{type:String,required:!0},onClick:Function,tmNode:{type:Object,required:!0},isEllipsisPlaceholder:Boolean},setup(e){const{props:t}=W(ae);return{menuProps:t,style:x(()=>{const{paddingLeft:o}=e;return{paddingLeft:o&&`${o}px`}}),iconStyle:x(()=>{const{maxIconSize:o,activeIconSize:n,iconMarginRight:c}=e;return{width:`${o}px`,height:`${o}px`,fontSize:`${n}px`,marginRight:`${c}px`}})}},render(){const{clsPrefix:e,tmNode:t,menuProps:{renderIcon:o,renderLabel:n,renderExtra:c,expandIcon:i}}=this,u=o?o(t.rawNode):te(this.icon);return(()=>{const h=we("7bb10afc6caf8fa4");return a(),g("div",{onClick:v=>{this.onClick?.(v)},role:"none",class:T([`${e}-menu-item-content`,{[`${e}-menu-item-content--selected`]:this.selected,[`${e}-menu-item-content--collapsed`]:this.collapsed,[`${e}-menu-item-content--child-active`]:this.childActive,[`${e}-menu-item-content--disabled`]:this.disabled,[`${e}-menu-item-content--hover`]:this.hover}]),style:U(this.style)},[R(()=>u&&(a(),g("div",{class:T(`${e}-menu-item-content__icon`),style:U(this.iconStyle),role:"none"},[R(()=>[u])],6))),y("div",{class:T(`${e}-menu-item-content-header`),role:"none"},[this.isEllipsisPlaceholder?(a(),g(D,{key:0},[R(()=>this.title)],64)):(a(),g(D,{key:1},[n?(a(),g(D,{key:0},[R(()=>n(t.rawNode))],64)):(a(),g(D,{key:1},[R(()=>te(this.title))],64))],64)),this.extra||c?(a(),g("span",{key:2,class:T(`${e}-menu-item-content-header__extra`)},[h[0]||(h[0]=R(" ",-1)),c?(a(),g(D,{key:0},[R(()=>c(t.rawNode))],64)):(a(),g(D,{key:1},[R(()=>te(this.extra))],64))],2)):R(()=>null)],2),this.showArrow?(a(),B(De,{key:0,ariaHidden:!0,class:T(`${e}-menu-item-content__arrow`),clsPrefix:e},{default:()=>i?i(t.rawNode):(a(),B(ot,{key:1}))},1032,["class","clsPrefix"])):R(()=>null)],14,tt)})()}});const se=8;function Ae(e){const t=W(ae),{props:o,mergedCollapsedRef:n}=t,c=W(Ye,null),i=W(Se,null),u=x(()=>o.mode==="horizontal"),h=x(()=>u.value?o.dropdownPlacement:"tmNodes"in e?"right-start":"right"),v=x(()=>Math.max(o.collapsedIconSize??o.iconSize,o.iconSize));return{dropdownPlacement:h,activeIconSize:x(()=>!u.value&&e.root&&n.value?o.collapsedIconSize??o.iconSize:o.iconSize),maxIconSize:v,paddingLeft:x(()=>{if(u.value)return;const{collapsedWidth:m,indent:_,rootIndent:A}=o,{root:I,isGroup:w}=e,C=A===void 0?_:A;return I?n.value?m/2-v.value/2:C:i&&typeof i.paddingLeftRef.value=="number"?n.value?m/2-v.value/2:_/2+i.paddingLeftRef.value:c&&typeof c.paddingLeftRef.value=="number"?(w?_/2:_)+c.paddingLeftRef.value:0}),iconMarginRight:x(()=>{const{collapsedWidth:m,indent:_,rootIndent:A}=o,{value:I}=v,{root:w}=e;return u.value||!w||!n.value?se:(A===void 0?_:A)+I+se-(m+I)/2}),NMenu:t,NSubmenu:c,NMenuOptionGroup:i}}const He={internalKey:{type:[String,Number],required:!0},root:Boolean,isGroup:Boolean,level:{type:Number,required:!0},title:[String,Function],extra:[String,Function]},Ze={...He,tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function},rt=ze(Ze),nt=$({name:"MenuOption",props:Ze,setup(e){const t=Ae(e),{NSubmenu:o,NMenu:n,NMenuOptionGroup:c}=t,{props:i,mergedClsPrefixRef:u,mergedCollapsedRef:h}=n,v=o?o.mergedDisabledRef:c?c.mergedDisabledRef:{value:!1},m=x(()=>v.value||e.disabled);function _(I){const{onClick:w}=e;w&&w(I)}function A(I){m.value||(n.doSelect(e.internalKey,e.tmNode.rawNode),_(I))}return{mergedClsPrefix:u,dropdownPlacement:t.dropdownPlacement,paddingLeft:t.paddingLeft,iconMarginRight:t.iconMarginRight,maxIconSize:t.maxIconSize,activeIconSize:t.activeIconSize,mergedTheme:n.mergedThemeRef,menuProps:i,dropdownEnabled:xe(()=>e.root&&h.value&&i.mode!=="horizontal"&&!m.value),selected:xe(()=>n.mergedValueRef.value===e.internalKey),mergedDisabled:m,handleClick:A}},render(){const{mergedClsPrefix:e,mergedTheme:t,tmNode:o,menuProps:{renderLabel:n,nodeProps:c}}=this,i=c?.(o.rawNode);return a(),g("div",X(i,{role:"menuitem",class:[`${e}-menu-item`,i?.class]}),[(a(),B(Ro,{theme:t.peers.Tooltip,themeOverrides:t.peerOverrides.Tooltip,trigger:"hover",placement:this.dropdownPlacement,disabled:!this.dropdownEnabled||this.title===void 0,internalExtraClass:["menu-tooltip"]},{default:()=>n?n(o.rawNode):te(this.title),trigger:()=>(a(),B(Xe,{tmNode:o,clsPrefix:e,paddingLeft:this.paddingLeft,iconMarginRight:this.iconMarginRight,maxIconSize:this.maxIconSize,activeIconSize:this.activeIconSize,selected:this.selected,title:this.title,extra:this.extra,disabled:this.mergedDisabled,icon:this.icon,onClick:this.handleClick},null,8,["tmNode","clsPrefix","paddingLeft","iconMarginRight","maxIconSize","activeIconSize","selected","title","extra","disabled","icon","onClick"]))},1032,["theme","themeOverrides","placement","disabled"]))],16)}}),Je={...He,tmNode:{type:Object,required:!0},tmNodes:{type:Array,required:!0}},lt=ze(Je),it=$({name:"MenuOptionGroup",props:Je,setup(e){const t=Ae(e),{NSubmenu:o}=t,n=x(()=>o?.mergedDisabledRef.value?!0:e.tmNode.disabled);ne(Se,{paddingLeftRef:t.paddingLeft,mergedDisabledRef:n});const{mergedClsPrefixRef:c,props:i}=W(ae);return function(){const{value:u}=c,h=t.paddingLeft.value,{nodeProps:v}=i,m=v?.(e.tmNode.rawNode);return(()=>{const _=we("45eca6a63be5028b");return a(),g("div",{class:T(`${u}-menu-item-group`),role:"group"},[y("div",X(m,{class:[`${u}-menu-item-group-title`,m?.class],style:[m?.style||"",h!==void 0?`padding-left: ${h}px;`:""]}),[R(()=>te(e.title)),e.extra?(a(),g(D,{key:0},[_[0]||(_[0]=R(" ",-1)),R(()=>te(e.extra))],64)):R(()=>null)],16),y("div",null,[R(()=>e.tmNodes.map(A=>Re(A,i)))])],2)})()}}}),at=["aria-expanded","id"],ct=["aria-expanded","id"],Qe={...He,rawNodes:{type:Array,default:()=>[]},tmNodes:{type:Array,default:()=>[]},tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function,domId:String,virtualChildActive:{type:Boolean,default:void 0},isEllipsisPlaceholder:Boolean},st=ze(Qe),be=$({name:"Submenu",props:Qe,setup(e){const t=Ae(e),{NMenu:o,NSubmenu:n}=t,{props:c,mergedCollapsedRef:i,mergedThemeRef:u}=o,h=x(()=>{const{disabled:I}=e;return n?.mergedDisabledRef.value||c.disabled?!0:I}),v=F(!1);ne(Ye,{paddingLeftRef:t.paddingLeft,mergedDisabledRef:h}),ne(Se,null);function m(){const{onClick:I}=e;I&&I()}function _(){h.value||(i.value||o.toggleExpand(e.internalKey),m())}function A(I){v.value=I}return{menuProps:c,mergedTheme:u,doSelect:o.doSelect,inverted:o.invertedRef,isHorizontal:o.isHorizontalRef,mergedClsPrefix:o.mergedClsPrefixRef,maxIconSize:t.maxIconSize,activeIconSize:t.activeIconSize,iconMarginRight:t.iconMarginRight,dropdownPlacement:t.dropdownPlacement,dropdownShow:v,paddingLeft:t.paddingLeft,mergedDisabled:h,mergedValue:o.mergedValueRef,childActive:xe(()=>e.virtualChildActive??o.activePathRef.value.includes(e.internalKey)),collapsed:x(()=>c.mode==="horizontal"?!1:i.value?!0:!o.mergedExpandedKeysRef.value.includes(e.internalKey)),dropdownEnabled:x(()=>!h.value&&(c.mode==="horizontal"||i.value)),handlePopoverShowChange:A,handleClick:_}},render(){const{mergedClsPrefix:e,menuProps:{renderIcon:t,renderLabel:o}}=this,n=()=>{const{isHorizontal:i,paddingLeft:u,collapsed:h,mergedDisabled:v,maxIconSize:m,activeIconSize:_,title:A,childActive:I,icon:w,handleClick:C,menuProps:{nodeProps:b},dropdownShow:P,iconMarginRight:M,tmNode:ee,mergedClsPrefix:G,isEllipsisPlaceholder:oe,extra:N}=this,H=b?.(ee.rawNode);return a(),g("div",X(H,{class:[`${G}-menu-item`,H?.class],role:"menuitem"}),[(a(),B(Xe,{tmNode:ee,paddingLeft:u,collapsed:h,disabled:v,iconMarginRight:M,maxIconSize:m,activeIconSize:_,title:A,extra:N,showArrow:!i,childActive:I,clsPrefix:G,icon:w,hover:P,onClick:C,isEllipsisPlaceholder:oe},null,8,["tmNode","paddingLeft","collapsed","disabled","iconMarginRight","maxIconSize","activeIconSize","title","extra","showArrow","childActive","clsPrefix","icon","hover","onClick","isEllipsisPlaceholder"]))],16)},c=()=>(a(),B(ho,null,{default:()=>{const{tmNodes:i,collapsed:u}=this;return u?null:(a(),g("div",{key:1,class:T(`${e}-submenu-children`),role:"menu"},[R(()=>i.map(h=>Re(h,this.menuProps)))],2))}},1024));return this.root?(a(),B(Ue,X({key:2,size:"large",trigger:"hover"},this.menuProps?.dropdownProps,{themeOverrides:this.mergedTheme.peerOverrides.Dropdown,theme:this.mergedTheme.peers.Dropdown,builtinThemeOverrides:{fontSizeLarge:"14px",optionIconSizeLarge:"18px"},value:this.mergedValue,disabled:!this.dropdownEnabled,placement:this.dropdownPlacement,keyField:this.menuProps.keyField,labelField:this.menuProps.labelField,childrenField:this.menuProps.childrenField,onUpdateShow:this.handlePopoverShowChange,options:this.rawNodes,onSelect:this.doSelect,inverted:this.inverted,renderIcon:t,renderLabel:o}),{default:()=>(a(),g("div",{class:T(`${e}-submenu`),role:"menu","aria-expanded":!this.collapsed,id:this.domId},[R(()=>n()),this.isHorizontal?R(()=>null):(a(),g(D,{key:1},[R(()=>c())],64))],10,at))},1040,["themeOverrides","theme","value","disabled","placement","keyField","labelField","childrenField","onUpdateShow","options","onSelect","inverted","renderIcon","renderLabel"])):(a(),g("div",{key:3,class:T(`${e}-submenu`),role:"menu","aria-expanded":!this.collapsed,id:this.domId},[R(()=>n()),R(()=>c())],10,ct))}});function ye(e){return e.type==="divider"||e.type==="render"}function dt(e){return e.type==="divider"}function Re(e,t){const{rawNode:o}=e,{show:n}=o;if(n===!1)return null;if(ye(o))return dt(o)?(a(),B(et,X({key:e.key},o.props),null,16)):null;const{labelField:c}=t,{key:i,level:u,isGroup:h}=e,v={...o,title:o.title||o[c],extra:o.titleExtra||o.extra,key:i,internalKey:i,level:u,root:u===0,isGroup:h};return e.children?e.isGroup?re(it,pe(v,lt,{tmNode:e,tmNodes:e.children,key:i})):re(be,pe(v,st,{key:i,rawNodes:o[t.childrenField],tmNodes:e.children,tmNode:e})):re(nt,pe(v,rt,{key:i,tmNode:e}))}const ut={...Y.props,options:{type:Array,default:()=>[]},collapsed:{type:Boolean,default:void 0},collapsedWidth:{type:Number,default:48},iconSize:{type:Number,default:20},collapsedIconSize:{type:Number,default:24},rootIndent:Number,indent:{type:Number,default:32},labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},disabledField:{type:String,default:"disabled"},defaultExpandAll:Boolean,defaultExpandedKeys:Array,expandedKeys:Array,value:[String,Number],defaultValue:{type:[String,Number],default:null},mode:{type:String,default:"vertical"},watchProps:{type:Array,default:void 0},disabled:Boolean,show:{type:Boolean,default:!0},inverted:Boolean,"onUpdate:expandedKeys":[Function,Array],onUpdateExpandedKeys:[Function,Array],onUpdateValue:[Function,Array],"onUpdate:value":[Function,Array],expandIcon:Function,renderIcon:Function,renderLabel:Function,renderExtra:Function,dropdownProps:Object,accordion:Boolean,nodeProps:Function,dropdownPlacement:{type:String,default:"bottom"},responsive:Boolean,items:Array,onOpenNamesChange:[Function,Array],onSelect:[Function,Array],onExpandedNamesChange:[Function,Array],expandedNames:Array,defaultExpandedNames:Array};var vt=$({name:"Menu",inheritAttrs:!1,props:ut,setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:o}=ue(e),n=Y("Menu","-menu",Qo,Oo,e,t),c=W(Ge,null),i=x(()=>{const{collapsed:p}=e;if(p!==void 0)return p;if(c){const{collapseModeRef:k,collapsedRef:r}=c;if(k.value==="width")return r.value??!1}return!1}),u=x(()=>{const{keyField:p,childrenField:k,disabledField:r}=e;return fe(e.items||e.options,{getIgnored(f){return ye(f)},getChildren(f){return f[k]},getDisabled(f){return f[r]},getKey(f){return f[p]??f.name}})}),h=x(()=>new Set(u.value.treeNodes.map(p=>p.key))),{watchProps:v}=e,m=F(null);v?.includes("defaultValue")?Te(()=>{m.value=e.defaultValue}):m.value=e.defaultValue;const _=de(e,"value"),A=Ce(_,m),I=F([]),w=()=>{I.value=e.defaultExpandAll?u.value.getNonLeafKeys():e.defaultExpandedNames||e.defaultExpandedKeys||u.value.getPath(A.value,{includeSelf:!1}).keyPath};v?.includes("defaultExpandedKeys")?Te(w):w();const C=Bo(e,["expandedNames","expandedKeys"]),b=Ce(C,I),P=x(()=>u.value.treeNodes),M=x(()=>u.value.getPath(A.value).keyPath);ne(ae,{props:e,mergedCollapsedRef:i,mergedThemeRef:n,mergedValueRef:A,mergedExpandedKeysRef:b,activePathRef:M,mergedClsPrefixRef:t,isHorizontalRef:x(()=>e.mode==="horizontal"),invertedRef:de(e,"inverted"),doSelect:ee,toggleExpand:oe});function ee(p,k){const{"onUpdate:value":r,onUpdateValue:f,onSelect:L}=e;f&&V(f,p,k),r&&V(r,p,k),L&&V(L,p,k),m.value=p}function G(p){const{"onUpdate:expandedKeys":k,onUpdateExpandedKeys:r,onExpandedNamesChange:f,onOpenNamesChange:L}=e;k&&V(k,p),r&&V(r,p),f&&V(f,p),L&&V(L,p),I.value=p}function oe(p){const k=Array.from(b.value),r=k.findIndex(f=>f===p);if(~r)k.splice(r,1);else{if(e.accordion&&h.value.has(p)){const f=k.findIndex(L=>h.value.has(L));f>-1&&k.splice(f,1)}k.push(p)}G(k)}const N=p=>{const k=u.value.getPath(p??A.value,{includeSelf:!1}).keyPath;if(!k.length)return;const r=Array.from(b.value),f=new Set([...r,...k]);e.accordion&&h.value.forEach(L=>{f.has(L)&&!k.includes(L)&&f.delete(L)}),G(Array.from(f))},H=x(()=>{const{inverted:p}=e,{common:{cubicBezierEaseInOut:k},self:r}=n.value,{borderRadius:f,borderColorHorizontal:L,fontSize:ao,itemHeight:co,dividerColor:so}=r,l={"--n-divider-color":so,"--n-bezier":k,"--n-font-size":ao,"--n-border-color-horizontal":L,"--n-border-radius":f,"--n-item-height":co};return p?(l["--n-group-text-color"]=r.groupTextColorInverted,l["--n-color"]=r.colorInverted,l["--n-item-text-color"]=r.itemTextColorInverted,l["--n-item-text-color-hover"]=r.itemTextColorHoverInverted,l["--n-item-text-color-active"]=r.itemTextColorActiveInverted,l["--n-item-text-color-child-active"]=r.itemTextColorChildActiveInverted,l["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveInverted,l["--n-item-text-color-active-hover"]=r.itemTextColorActiveHoverInverted,l["--n-item-icon-color"]=r.itemIconColorInverted,l["--n-item-icon-color-hover"]=r.itemIconColorHoverInverted,l["--n-item-icon-color-active"]=r.itemIconColorActiveInverted,l["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHoverInverted,l["--n-item-icon-color-child-active"]=r.itemIconColorChildActiveInverted,l["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHoverInverted,l["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsedInverted,l["--n-item-text-color-horizontal"]=r.itemTextColorHorizontalInverted,l["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontalInverted,l["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontalInverted,l["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontalInverted,l["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontalInverted,l["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontalInverted,l["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontalInverted,l["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontalInverted,l["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontalInverted,l["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontalInverted,l["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontalInverted,l["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontalInverted,l["--n-arrow-color"]=r.arrowColorInverted,l["--n-arrow-color-hover"]=r.arrowColorHoverInverted,l["--n-arrow-color-active"]=r.arrowColorActiveInverted,l["--n-arrow-color-active-hover"]=r.arrowColorActiveHoverInverted,l["--n-arrow-color-child-active"]=r.arrowColorChildActiveInverted,l["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHoverInverted,l["--n-item-color-hover"]=r.itemColorHoverInverted,l["--n-item-color-active"]=r.itemColorActiveInverted,l["--n-item-color-active-hover"]=r.itemColorActiveHoverInverted,l["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsedInverted):(l["--n-group-text-color"]=r.groupTextColor,l["--n-color"]=r.color,l["--n-item-text-color"]=r.itemTextColor,l["--n-item-text-color-hover"]=r.itemTextColorHover,l["--n-item-text-color-active"]=r.itemTextColorActive,l["--n-item-text-color-child-active"]=r.itemTextColorChildActive,l["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveHover,l["--n-item-text-color-active-hover"]=r.itemTextColorActiveHover,l["--n-item-icon-color"]=r.itemIconColor,l["--n-item-icon-color-hover"]=r.itemIconColorHover,l["--n-item-icon-color-active"]=r.itemIconColorActive,l["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHover,l["--n-item-icon-color-child-active"]=r.itemIconColorChildActive,l["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHover,l["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsed,l["--n-item-text-color-horizontal"]=r.itemTextColorHorizontal,l["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontal,l["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontal,l["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontal,l["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontal,l["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontal,l["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontal,l["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontal,l["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontal,l["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontal,l["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontal,l["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontal,l["--n-arrow-color"]=r.arrowColor,l["--n-arrow-color-hover"]=r.arrowColorHover,l["--n-arrow-color-active"]=r.arrowColorActive,l["--n-arrow-color-active-hover"]=r.arrowColorActiveHover,l["--n-arrow-color-child-active"]=r.arrowColorChildActive,l["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHover,l["--n-item-color-hover"]=r.itemColorHover,l["--n-item-color-active"]=r.itemColorActive,l["--n-item-color-active-hover"]=r.itemColorActiveHover,l["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsed),l}),O=o?ve("menu",x(()=>e.inverted?"a":"b"),H,e):void 0,q=po(),j=F(null),he=F(null);let E=!0;const _e=()=>{E?E=!1:j.value?.sync({showAllItemsBeforeCalculate:!0})};function eo(){return document.getElementById(q)}const ce=F(-1);function oo(p){ce.value=e.options.length-p}function to(p){p||(ce.value=-1)}const ro=x(()=>{const p=ce.value;return{children:p===-1?[]:e.options.slice(p)}}),no=x(()=>{const{childrenField:p,disabledField:k,keyField:r}=e;return fe([ro.value],{getIgnored(f){return ye(f)},getChildren(f){return f[p]},getDisabled(f){return f[k]},getKey(f){return f[r]??f.name}})}),lo=x(()=>fe([{}]).treeNodes[0]);function io(){if(ce.value===-1)return a(),B(be,{root:!0,level:0,key:"__ellpisisGroupPlaceholder__",internalKey:"__ellpisisGroupPlaceholder__",title:"···",tmNode:lo.value,domId:q,isEllipsisPlaceholder:!0},null,8,["tmNode","domId"]);const p=no.value.treeNodes[0],k=M.value,r=!!p.children?.some(f=>k.includes(f.key));return a(),B(be,{level:0,root:!0,key:"__ellpisisGroup__",internalKey:"__ellpisisGroup__",title:"···",virtualChildActive:r,tmNode:p,domId:q,rawNodes:p.rawNode.children||[],tmNodes:p.children||[],isEllipsisPlaceholder:!0},null,8,["virtualChildActive","tmNode","domId","rawNodes","tmNodes"])}return{mergedClsPrefix:t,controlledExpandedKeys:C,uncontrolledExpanededKeys:I,mergedExpandedKeys:b,uncontrolledValue:m,mergedValue:A,activePath:M,tmNodes:P,mergedTheme:n,mergedCollapsed:i,cssVars:o?void 0:H,themeClass:O?.themeClass,overflowRef:j,counterRef:he,updateCounter:()=>{},onResize:_e,onUpdateOverflow:to,onUpdateCount:oo,renderCounter:io,getCounter:eo,onRender:O?.onRender,showOption:N,deriveResponsiveState:_e}},render(){const{mergedClsPrefix:e,mode:t,themeClass:o,onRender:n}=this;n?.();const c=()=>this.tmNodes.map(h=>Re(h,this.$props)),i=t==="horizontal"&&this.responsive,u=()=>re("div",X(this.$attrs,{role:t==="horizontal"?"menubar":"menu",class:[`${e}-menu`,o,`${e}-menu--${t}`,i&&`${e}-menu--responsive`,this.mergedCollapsed&&`${e}-menu--collapsed`],style:this.cssVars}),i?(a(),B(_o,{key:2,ref:"overflowRef",onUpdateOverflow:this.onUpdateOverflow,getCounter:this.getCounter,onUpdateCount:this.onUpdateCount,updateCounter:this.updateCounter,style:{width:"100%",display:"flex",overflow:"hidden"}},{default:c,counter:this.renderCounter},1032,["onUpdateOverflow","getCounter","onUpdateCount","updateCounter"])):c());return i?(a(),B(mo,{key:3,onResize:this.onResize},{default:u},1032,["onResize"])):u()}});const ht={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},mt=$({name:"AppsOutline",render:function(t,o){return a(),g("svg",ht,o[0]||(o[0]=[fo('<rect x="64" y="64" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="216" y="64" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="368" y="64" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="64" y="216" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="216" y="216" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="368" y="216" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="64" y="368" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="216" y="368" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect><rect x="368" y="368" width="80" height="80" rx="40" ry="40" fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32"></rect>',9)]))}}),pt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},ft=$({name:"DocumentTextOutline",render:function(t,o){return a(),g("svg",pt,o[0]||(o[0]=[y("path",{d:"M416 221.25V416a48 48 0 0 1-48 48H144a48 48 0 0 1-48-48V96a48 48 0 0 1 48-48h98.75a32 32 0 0 1 22.62 9.37l141.26 141.26a32 32 0 0 1 9.37 22.62z",fill:"none",stroke:"currentColor","stroke-linejoin":"round","stroke-width":"32"},null,-1),y("path",{d:"M256 56v120a32 32 0 0 0 32 32h120",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),y("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M176 288h160"},null,-1),y("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M176 368h160"},null,-1)]))}}),gt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Oe=$({name:"GridOutline",render:function(t,o){return a(),g("svg",gt,o[0]||(o[0]=[y("rect",{x:"48",y:"48",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),y("rect",{x:"288",y:"48",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),y("rect",{x:"48",y:"288",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),y("rect",{x:"288",y:"288",width:"176",height:"176",rx:"20",ry:"20",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1)]))}}),xt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},Ee=$({name:"KeyOutline",render:function(t,o){return a(),g("svg",xt,o[0]||(o[0]=[y("path",{d:"M218.1 167.17c0 13 0 25.6 4.1 37.4c-43.1 50.6-156.9 184.3-167.5 194.5a20.17 20.17 0 0 0-6.7 15c0 8.5 5.2 16.7 9.6 21.3c6.6 6.9 34.8 33 40 28c15.4-15 18.5-19 24.8-25.2c9.5-9.3-1-28.3 2.3-36s6.8-9.2 12.5-10.4s15.8 2.9 23.7 3c8.3.1 12.8-3.4 19-9.2c5-4.6 8.6-8.9 8.7-15.6c.2-9-12.8-20.9-3.1-30.4s23.7 6.2 34 5s22.8-15.5 24.1-21.6s-11.7-21.8-9.7-30.7c.7-3 6.8-10 11.4-11s25 6.9 29.6 5.9c5.6-1.2 12.1-7.1 17.4-10.4c15.5 6.7 29.6 9.4 47.7 9.4c68.5 0 124-53.4 124-119.2S408.5 48 340 48s-121.9 53.37-121.9 119.17zM400 144a32 32 0 1 1-32-32a32 32 0 0 1 32 32z",fill:"none",stroke:"currentColor","stroke-linejoin":"round","stroke-width":"32"},null,-1)]))}}),Ct={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},bt=$({name:"LinkOutline",render:function(t,o){return a(),g("svg",Ct,o[0]||(o[0]=[y("path",{d:"M208 352h-64a96 96 0 0 1 0-192h64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),y("path",{d:"M304 160h64a96 96 0 0 1 0 192h-64",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36"},null,-1),y("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"36",d:"M163.29 256h187.42"},null,-1)]))}}),yt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},wt=$({name:"LogOutOutline",render:function(t,o){return a(),g("svg",yt,o[0]||(o[0]=[y("path",{d:"M304 336v40a40 40 0 0 1-40 40H104a40 40 0 0 1-40-40V136a40 40 0 0 1 40-40h152c22.09 0 48 17.91 48 40v40",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),y("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M368 336l80-80l-80-80"},null,-1),y("path",{fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32",d:"M176 256h256"},null,-1)]))}}),zt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},It=$({name:"PricetagsOutline",render:function(t,o){return a(),g("svg",zt,o[0]||(o[0]=[y("path",{d:"M403.29 32H280.36a14.46 14.46 0 0 0-10.2 4.2L24.4 281.9a28.85 28.85 0 0 0 0 40.7l117 117a28.86 28.86 0 0 0 40.71 0L427.8 194a14.46 14.46 0 0 0 4.2-10.2v-123A28.66 28.66 0 0 0 403.29 32z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1),y("path",{d:"M352 144a32 32 0 1 1 32-32a32 32 0 0 1-32 32z",fill:"currentColor"},null,-1),y("path",{d:"M230 480l262-262a13.81 13.81 0 0 0 4-10V80",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1)]))}}),kt={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 512 512"},St=$({name:"SettingsOutline",render:function(t,o){return a(),g("svg",kt,o[0]||(o[0]=[y("path",{d:"M262.29 192.31a64 64 0 1 0 57.4 57.4a64.13 64.13 0 0 0-57.4-57.4zM416.39 256a154.34 154.34 0 0 1-1.53 20.79l45.21 35.46a10.81 10.81 0 0 1 2.45 13.75l-42.77 74a10.81 10.81 0 0 1-13.14 4.59l-44.9-18.08a16.11 16.11 0 0 0-15.17 1.75A164.48 164.48 0 0 1 325 400.8a15.94 15.94 0 0 0-8.82 12.14l-6.73 47.89a11.08 11.08 0 0 1-10.68 9.17h-85.54a11.11 11.11 0 0 1-10.69-8.87l-6.72-47.82a16.07 16.07 0 0 0-9-12.22a155.3 155.3 0 0 1-21.46-12.57a16 16 0 0 0-15.11-1.71l-44.89 18.07a10.81 10.81 0 0 1-13.14-4.58l-42.77-74a10.8 10.8 0 0 1 2.45-13.75l38.21-30a16.05 16.05 0 0 0 6-14.08c-.36-4.17-.58-8.33-.58-12.5s.21-8.27.58-12.35a16 16 0 0 0-6.07-13.94l-38.19-30A10.81 10.81 0 0 1 49.48 186l42.77-74a10.81 10.81 0 0 1 13.14-4.59l44.9 18.08a16.11 16.11 0 0 0 15.17-1.75A164.48 164.48 0 0 1 187 111.2a15.94 15.94 0 0 0 8.82-12.14l6.73-47.89A11.08 11.08 0 0 1 213.23 42h85.54a11.11 11.11 0 0 1 10.69 8.87l6.72 47.82a16.07 16.07 0 0 0 9 12.22a155.3 155.3 0 0 1 21.46 12.57a16 16 0 0 0 15.11 1.71l44.89-18.07a10.81 10.81 0 0 1 13.14 4.58l42.77 74a10.8 10.8 0 0 1-2.45 13.75l-38.21 30a16.05 16.05 0 0 0-6.05 14.08c.33 4.14.55 8.3.55 12.47z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"},null,-1)]))}}),At={class:"brand"},Ht={class:"brand-text"},Rt={class:"header-title"},_t={class:"user pressable",type:"button"},Pt=$({__name:"AdminLayout",setup(e){const t=zo(),o=ko(),n=go(),c=To(),i=F(!1),u={dashboard:Oe,site:bt,category:It,config:St,about_page:ft,modify_password:Ee};function h(w){const C=w.split("/").filter(Boolean).pop()??"";return u[C]??mt}function v(w){return()=>re(Po,null,{default:()=>re(w)})}const m=x(()=>{const w=[{label:"仪表盘",key:"/admin/dashboard",icon:v(Oe)}];for(const C of n.menu.filter(b=>!b.pid)){const b=n.menu.filter(P=>P.pid===C.id);if(b.length===0){C.link&&w.push({label:C.name,key:C.link,icon:v(h(C.link))});continue}w.push({label:C.name,key:`group-${C.id}`,icon:v(h(C.link)),children:b.map(P=>({label:P.name,key:P.link}))})}return w}),_=[{label:"修改密码",key:"modify_password",icon:v(Ee)},{label:"退出登录",key:"logout",icon:v(wt)}];function A(w){w!==t.path&&o.push(w)}async function I(w){if(w==="modify_password"){o.push("/admin/modify_password");return}try{await n.logout(),o.replace("/login")}catch(C){c.error(Io(C))}}return(w,C)=>(a(),B(K(Be),{"has-sider":"",class:"layout"},{default:Z(()=>[J(K(Jo),{bordered:"","collapse-mode":"width","collapsed-width":64,width:208,collapsed:i.value,"show-trigger":"bar",onCollapse:C[0]||(C[0]=b=>i.value=!0),onExpand:C[1]||(C[1]=b=>i.value=!1)},{default:Z(()=>[y("div",At,[C[2]||(C[2]=y("span",{class:"brand-dot"},null,-1)),xo(y("span",Ht,"WebStack",512),[[Co,!i.value]])]),J(K(vt),{value:K(t).path,options:m.value,collapsed:i.value,"collapsed-width":64,"collapsed-icon-size":20,"onUpdate:value":A},null,8,["value","options","collapsed"])]),_:1},8,["collapsed"]),J(K(Be),null,{default:Z(()=>[J(K(Do),{bordered:"",class:"header"},{default:Z(()=>[y("span",Rt,Ne(K(t).meta.title??""),1),J(K(Ue),{options:_,trigger:"click",onSelect:I},{default:Z(()=>[y("button",_t,Ne(K(n).username||"管理员"),1)]),_:1})]),_:1}),J(K(Ko),{class:"content","content-style":"padding: 16px 20px 24px;"},{default:Z(()=>[J(K(bo),null,{default:Z(({Component:b})=>[(a(),B(yo,null,[(a(),B(wo(b)))],1024))]),_:1})]),_:1})]),_:1})]),_:1}))}}),Et=No(Pt,[["__scopeId","data-v-008db48c"]]);export{Et as default};
