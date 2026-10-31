import{b6 as Le,b7 as Me,ap as We,b8 as Ee,r as je,ar as oe,ao as he,as as ne,d as _,V as Ve,a as q,m as Ce,n as ve,p as Ue,q as qe,y as Q,_ as E,G as me,J as m,k as O,v as k,x as R,M as A,aC as Pe,F as J,I as ke,t as be,o as s,c as w,A as N,D as S,C as P,W as re,O as F,aA as Re,au as Ge,P as W,b9 as Ne,Y as te,K as G,ba as Je,bb as Qe,bc as Xe,al as Ye,T as ge,L as B,$ as Ze,Q as pe,R as z,ac as M}from"./index-Dx1zj60b.js";import{p as Ie,a as ye,P as _e,b as eo,d as oo,B as no,h as xe,r as to,c as ro}from"./create-Bo10zms7.js";import{f as io}from"./_plugin-vue_export-helper-CtyYQGco.js";import{a as ao}from"./use-message-U980NtEK.js";function lo(e={},o){const t=je({ctrl:!1,command:!1,win:!1,shift:!1,tab:!1}),{keydown:n,keyup:r}=e,a=l=>{switch(l.key){case"Control":t.ctrl=!0;break;case"Meta":t.command=!0,t.win=!0;break;case"Shift":t.shift=!0;break;case"Tab":t.tab=!0;break}n!==void 0&&Object.keys(n).forEach(y=>{if(y!==l.key)return;const f=n[y];if(typeof f=="function")f(l);else{const{stop:C=!1,prevent:g=!1}=f;C&&l.stopPropagation(),g&&l.preventDefault(),f.handler(l)}})},d=l=>{switch(l.key){case"Control":t.ctrl=!1;break;case"Meta":t.command=!1,t.win=!1;break;case"Shift":t.shift=!1;break;case"Tab":t.tab=!1;break}r!==void 0&&Object.keys(r).forEach(y=>{if(y!==l.key)return;const f=r[y];if(typeof f=="function")f(l);else{const{stop:C=!1,prevent:g=!1}=f;C&&l.stopPropagation(),g&&l.preventDefault(),f.handler(l)}})},c=()=>{(o===void 0||o.value)&&(oe("keydown",document,a),oe("keyup",document,d)),o!==void 0&&he(o,l=>{l?(oe("keydown",document,a),oe("keyup",document,d)):(ne("keydown",document,a),ne("keyup",document,d))})};return Le()?(Me(c),We(()=>{(o===void 0||o.value)&&(ne("keydown",document,a),ne("keyup",document,d))})):c(),Ee(t)}var so=_({name:"ChevronRight",render(){return(()=>{const e=Ve("6ab04425f4fcb756");return e[0]||(e[0]=q("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},[q("path",{d:"M5.64645 3.14645C5.45118 3.34171 5.45118 3.65829 5.64645 3.85355L9.79289 8L5.64645 12.1464C5.45118 12.3417 5.45118 12.6583 5.64645 12.8536C5.84171 13.0488 6.15829 13.0488 6.35355 12.8536L10.8536 8.35355C11.0488 8.15829 11.0488 7.84171 10.8536 7.64645L6.35355 3.14645C6.15829 2.95118 5.84171 2.95118 5.64645 3.14645Z",fill:"currentColor"})],-1))})()}});function co(e){return o=>{o?e.value=o.$el:e.value=null}}var uo={padding:"4px 0",optionIconSizeSmall:"14px",optionIconSizeMedium:"16px",optionIconSizeLarge:"16px",optionIconSizeHuge:"18px",optionSuffixWidthSmall:"14px",optionSuffixWidthMedium:"14px",optionSuffixWidthLarge:"16px",optionSuffixWidthHuge:"16px",optionIconSuffixWidthSmall:"32px",optionIconSuffixWidthMedium:"32px",optionIconSuffixWidthLarge:"36px",optionIconSuffixWidthHuge:"36px",optionPrefixWidthSmall:"14px",optionPrefixWidthMedium:"14px",optionPrefixWidthLarge:"16px",optionPrefixWidthHuge:"16px",optionIconPrefixWidthSmall:"36px",optionIconPrefixWidthMedium:"36px",optionIconPrefixWidthLarge:"40px",optionIconPrefixWidthHuge:"40px"};function po(e){const{primaryColor:o,textColor2:t,dividerColor:n,hoverColor:r,popoverColor:a,invertedColor:d,borderRadius:c,fontSizeSmall:l,fontSizeMedium:y,fontSizeLarge:f,fontSizeHuge:C,heightSmall:g,heightMedium:K,heightLarge:T,heightHuge:I,textColor3:$,opacityDisabled:D}=e;return{...uo,optionHeightSmall:g,optionHeightMedium:K,optionHeightLarge:T,optionHeightHuge:I,borderRadius:c,fontSizeSmall:l,fontSizeMedium:y,fontSizeLarge:f,fontSizeHuge:C,optionTextColor:t,optionTextColorHover:t,optionTextColorActive:o,optionTextColorChildActive:o,color:a,dividerColor:n,suffixColor:t,prefixColor:t,optionColorHover:r,optionColorActive:Ue(o,{alpha:.1}),groupHeaderTextColor:$,optionTextColorInverted:"#BBB",optionTextColorHoverInverted:"#FFF",optionTextColorActiveInverted:"#FFF",optionTextColorChildActiveInverted:"#FFF",colorInverted:d,dividerColorInverted:"#BBB",suffixColorInverted:"#BBB",prefixColorInverted:"#BBB",optionColorHoverInverted:o,optionColorActiveInverted:o,groupHeaderTextColorInverted:"#AAA",optionOpacityDisabled:D}}const fo=Ce({name:"Dropdown",common:ve,peers:{Popover:Ie},self:po});var ho={padding:"8px 14px"};function vo(e){const{borderRadius:o,boxShadow2:t,baseColor:n}=e;return{...ho,borderRadius:o,boxShadow:t,color:qe(n,"rgba(0, 0, 0, .85)"),textColor:n}}const mo=Ce({name:"Tooltip",common:ve,peers:{Popover:Ie},self:vo}),bo={...ye,...Q.props};var Oo=_({name:"Tooltip",props:bo,slots:Object,__popover__:!0,setup(e){const{mergedClsPrefixRef:o}=me(e),t=Q("Tooltip","-tooltip",void 0,mo,e,o),n=O(null);return{syncPosition(){n.value.syncPosition()},setShow(r){n.value.setShow(r)},popoverRef:n,mergedTheme:t,popoverThemeOverrides:m(()=>t.value.self)}},render(){const{mergedTheme:e,internalExtraClass:o}=this;return E(_e,{...this.$props,theme:e.peers.Popover,themeOverrides:e.peerOverrides.Popover,builtinThemeOverrides:this.popoverThemeOverrides,internalExtraClass:o.concat("tooltip"),ref:"popoverRef"},this.$slots)}});function yo(e){const{textColorBase:o,opacity1:t,opacity2:n,opacity3:r,opacity4:a,opacity5:d}=e;return{color:o,opacity1Depth:t,opacity2Depth:n,opacity3Depth:r,opacity4Depth:a,opacity5Depth:d}}const wo={common:ve,self:yo};var go=k("icon",`
 height: 1em;
 width: 1em;
 line-height: 1em;
 text-align: center;
 display: inline-block;
 position: relative;
 fill: currentColor;
`,[R("color-transition",{transition:"color .3s var(--n-bezier)"}),R("depth",{color:"var(--n-color)"},[A("svg",{opacity:"var(--n-opacity)",transition:"opacity .3s var(--n-bezier)"})]),A("svg",{height:"1em",width:"1em"})]);const xo={...Q.props,depth:[String,Number],size:[Number,String],color:String,component:[Object,Function]},So=_({_n_icon__:!0,name:"Icon",inheritAttrs:!1,props:xo,setup(e){const{mergedClsPrefixRef:o,inlineThemeDisabled:t}=me(e),n=Q("Icon","-icon",go,wo,e,o),r=m(()=>{const{depth:d}=e,{common:{cubicBezierEaseInOut:c},self:l}=n.value;if(d!==void 0){const{color:y,[`opacity${d}Depth`]:f}=l;return{"--n-bezier":c,"--n-color":y,"--n-opacity":f}}return{"--n-bezier":c,"--n-color":"","--n-opacity":""}}),a=t?ke("icon",m(()=>`${e.depth||"d"}`),r,e):void 0;return{mergedClsPrefix:o,mergedStyle:m(()=>{const{size:d,color:c}=e;return{fontSize:io(d),color:c}}),cssVars:t?void 0:r,themeClass:a?.themeClass,onRender:a?.onRender}},render(){const{$parent:e,depth:o,mergedClsPrefix:t,component:n,onRender:r,themeClass:a}=this;return e?.$options?._n_icon__&&Pe("icon","don't wrap `n-icon` inside `n-icon`"),r?.(),E("i",J(this.$attrs,{role:"img",class:[`${t}-icon`,a,{[`${t}-icon--depth`]:o,[`${t}-icon--color-transition`]:o!==void 0}],style:[this.cssVars,this.mergedStyle]}),n?E(n):this.$slots.default?.())}}),we=be("n-dropdown-menu"),ie=be("n-dropdown"),Se=be("n-dropdown-option");var ze=_({name:"DropdownDivider",props:{clsPrefix:{type:String,required:!0}},render(){return s(),w("div",{class:N(`${this.clsPrefix}-dropdown-divider`)},null,2)}});function fe(e,o){return e.type==="submenu"||e.type===void 0&&e[o]!==void 0}function Co(e){return e.type==="group"}function Ke(e){return e.type==="divider"}function Po(e){return e.type==="render"}function ko(e,o,t){const n=O(e.value);let r=null;return he(e,a=>{r!==null&&window.clearTimeout(r),a===!0?t&&!t.value?n.value=!0:r=window.setTimeout(()=>{n.value=!0},o):n.value=!1}),n}var Te=_({name:"DropdownOption",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0},parentKey:{type:[String,Number],default:null},placement:{type:String,default:"right-start"},props:Object,scrollable:Boolean},setup(e){const o=W(ie),{hoverKeyRef:t,keyboardKeyRef:n,lastToggledSubmenuKeyRef:r,pendingKeyPathRef:a,activeKeyPathRef:d,animatedRef:c,mergedShowRef:l,renderLabelRef:y,renderIconRef:f,labelFieldRef:C,childrenFieldRef:g,renderOptionRef:K,nodePropsRef:T,menuPropsRef:I}=o,$=W(Se,null),D=W(we),j=W(Ne),Z=m(()=>e.tmNode.rawNode),X=m(()=>{const{value:i}=g;return fe(e.tmNode.rawNode,i)}),ae=m(()=>{const{disabled:i}=e.tmNode;return i}),de=m(()=>{if(!X.value)return!1;const{key:i,disabled:u}=e.tmNode;if(u)return!1;const{value:x}=t,{value:H}=n,{value:ue}=r,{value:L}=a;return x!==null?L.includes(i):H!==null?L.includes(i)&&L[L.length-1]!==i:ue!==null?L.includes(i):!1}),le=m(()=>n.value===null&&!c.value),se=ko(de,300,le),ce=m(()=>!!$?.enteringSubmenuRef.value),V=O(!1);G(Se,{enteringSubmenuRef:V});function U(){V.value=!0}function ee(){V.value=!1}function Y(){const{parentKey:i,tmNode:u}=e;u.disabled||l.value&&(r.value=i,n.value=null,t.value=u.key)}function p(){const{tmNode:i}=e;i.disabled||l.value&&t.value!==i.key&&Y()}function h(i){if(e.tmNode.disabled||!l.value)return;const{relatedTarget:u}=i;u&&!xe({target:u},"dropdownOption")&&!xe({target:u},"scrollbarRail")&&(t.value=null)}function b(){const{value:i}=X,{tmNode:u}=e;l.value&&!i&&!u.disabled&&(o.doSelect(u.key,u.rawNode),o.doUpdateShow(!1))}return{labelField:C,renderLabel:y,renderIcon:f,siblingHasIcon:D.showIconRef,siblingHasSubmenu:D.hasSubmenuRef,menuProps:I,popoverBody:j,animated:c,mergedShowSubmenu:m(()=>se.value&&!ce.value),rawNode:Z,hasSubmenu:X,pending:te(()=>{const{value:i}=a,{key:u}=e.tmNode;return i.includes(u)}),childActive:te(()=>{const{value:i}=d,{key:u}=e.tmNode,x=i.findIndex(H=>u===H);return x===-1?!1:x<i.length-1}),active:te(()=>{const{value:i}=d,{key:u}=e.tmNode,x=i.findIndex(H=>u===H);return x===-1?!1:x===i.length-1}),mergedDisabled:ae,renderOption:K,nodeProps:T,handleClick:b,handleMouseMove:p,handleMouseEnter:Y,handleMouseLeave:h,handleSubmenuBeforeEnter:U,handleSubmenuAfterEnter:ee}},render(){const{animated:e,rawNode:o,mergedShowSubmenu:t,clsPrefix:n,siblingHasIcon:r,siblingHasSubmenu:a,renderLabel:d,renderIcon:c,renderOption:l,nodeProps:y,props:f,scrollable:C}=this;let g=null;if(t){const $=this.menuProps?.(o,o.children);g=(D=>(s(),S($e,J({key:1},$,{clsPrefix:n,scrollable:this.scrollable,tmNodes:this.tmNode.children,parentKey:this.tmNode.key}),null,16,["clsPrefix","scrollable","tmNodes","parentKey"])))()}const K={class:[`${n}-dropdown-option-body`,this.pending&&`${n}-dropdown-option-body--pending`,this.active&&`${n}-dropdown-option-body--active`,this.childActive&&`${n}-dropdown-option-body--child-active`,this.mergedDisabled&&`${n}-dropdown-option-body--disabled`],onMousemove:this.handleMouseMove,onMouseenter:this.handleMouseEnter,onMouseleave:this.handleMouseLeave,onClick:this.handleClick},T=y?.(o),I=(s(),w("div",J({class:[`${n}-dropdown-option`,T?.class],"data-dropdown-option":!0},T),[P(()=>E("div",J(K,f),[(s(),w("div",{class:N([`${n}-dropdown-option-body__prefix`,r&&`${n}-dropdown-option-body__prefix--show-icon`])},[P(()=>[c?c(o):re(o.icon)])],2)),(s(),w("div",{"data-dropdown-option":!0,class:N(`${n}-dropdown-option-body__label`)},[d?(s(),w(F,{key:0},[P(()=>d(o))],64)):(s(),w(F,{key:1},[P(()=>re(o[this.labelField]??o.title))],64))],2)),(s(),w("div",{"data-dropdown-option":!0,class:N([`${n}-dropdown-option-body__suffix`,a&&`${n}-dropdown-option-body__suffix--has-submenu`])},[this.hasSubmenu?(s(),S(So,{key:0},{_:1,default:Re(()=>(s(),S(so)))})):P(()=>null)],2))])),this.hasSubmenu?(s(),S(no,{key:0},{default:()=>[(s(),S(eo,null,{default:()=>(s(),w("div",{class:N(`${n}-dropdown-offset-container`)},[(s(),S(oo,{show:this.mergedShowSubmenu,placement:this.placement,to:C&&this.popoverBody||void 0,teleportDisabled:!C},{default:()=>(s(),w("div",{class:N(`${n}-dropdown-menu-wrapper`)},[e?(s(),S(Ge,{key:0,onBeforeEnter:this.handleSubmenuBeforeEnter,onAfterEnter:this.handleSubmenuAfterEnter,name:"fade-in-scale-up-transition",appear:!0},{default:()=>g},1032,["onBeforeEnter","onAfterEnter"])):(s(),w(F,{key:1},[P(()=>g)],64))],2))},1032,["show","placement","to","teleportDisabled"]))],2))},1024))]},1024)):P(()=>null)],16));return l?l({node:I,option:o}):I}}),Ro=_({name:"DropdownGroupHeader",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(){const{showIconRef:e,hasSubmenuRef:o}=W(we),{renderLabelRef:t,labelFieldRef:n,nodePropsRef:r,renderOptionRef:a}=W(ie);return{labelField:n,showIcon:e,hasSubmenu:o,renderLabel:t,nodeProps:r,renderOption:a}},render(){const{clsPrefix:e,hasSubmenu:o,showIcon:t,nodeProps:n,renderLabel:r,renderOption:a}=this,{rawNode:d}=this.tmNode,c=(s(),w("div",J({class:`${e}-dropdown-option`},n?.(d)),[q("div",{class:N(`${e}-dropdown-option-body ${e}-dropdown-option-body--group`)},[q("div",{"data-dropdown-option":!0,class:N([`${e}-dropdown-option-body__prefix`,t&&`${e}-dropdown-option-body__prefix--show-icon`])},[P(()=>re(d.icon))],2),q("div",{class:N(`${e}-dropdown-option-body__label`),"data-dropdown-option":!0},[r?(s(),w(F,{key:0},[P(()=>r(d))],64)):(s(),w(F,{key:1},[P(()=>re(d.title??d[this.labelField]))],64))],2),q("div",{class:N([`${e}-dropdown-option-body__suffix`,o&&`${e}-dropdown-option-body__suffix--has-submenu`]),"data-dropdown-option":!0},null,2)],2)],16));return a?a({node:c,option:d}):c}}),No=_({name:"NDropdownGroup",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0},parentKey:{type:[String,Number],default:null}},render(){const{tmNode:e,parentKey:o,clsPrefix:t}=this,{children:n}=e;return s(),w(F,null,[(s(),S(Ro,{clsPrefix:t,tmNode:e,key:e.key},null,8,["clsPrefix","tmNode"])),P(()=>n?.map(r=>{const{rawNode:a}=r;return a.show===!1?null:Ke(a)?E(ze,{clsPrefix:t,key:r.key}):r.isGroup?(Pe("dropdown","`group` node is not allowed to be put in `group` node."),null):(s(),S(Te,{clsPrefix:t,tmNode:r,parentKey:o,key:r.key},null,8,["clsPrefix","tmNode","parentKey"]))}))],64)}}),Io=_({name:"DropdownRenderOption",props:{tmNode:{type:Object,required:!0}},render(){const{rawNode:{render:e,props:o}}=this.tmNode;return E("div",o,[e?.()])}}),$e=_({name:"DropdownMenu",props:{scrollable:Boolean,showArrow:Boolean,arrowStyle:[String,Object],clsPrefix:{type:String,required:!0},tmNodes:{type:Array,default:()=>[]},parentKey:{type:[String,Number],default:null}},setup(e){const{renderIconRef:o,childrenFieldRef:t}=W(ie);G(we,{showIconRef:m(()=>{const r=o.value;return e.tmNodes.some(a=>{if(a.isGroup)return a.children?.some(({rawNode:c})=>r?r(c):c.icon);const{rawNode:d}=a;return r?r(d):d.icon})}),hasSubmenuRef:m(()=>{const{value:r}=t;return e.tmNodes.some(a=>{if(a.isGroup)return a.children?.some(({rawNode:c})=>fe(c,r));const{rawNode:d}=a;return fe(d,r)})})});const n=O(null);return G(Qe,null),G(Xe,null),G(Ne,n),{bodyRef:n}},render(){const{parentKey:e,clsPrefix:o,scrollable:t}=this,n=this.tmNodes.map(r=>{const{rawNode:a}=r;return a.show===!1?null:Po(a)?(s(),S(Io,{tmNode:r,key:r.key},null,8,["tmNode"])):Ke(a)?(s(),S(ze,{clsPrefix:o,key:r.key},null,8,["clsPrefix"])):Co(a)?(s(),S(No,{clsPrefix:o,tmNode:r,parentKey:e,key:r.key},null,8,["clsPrefix","tmNode","parentKey"])):(s(),S(Te,{clsPrefix:o,tmNode:r,parentKey:e,key:r.key,props:a.props,scrollable:t},null,8,["clsPrefix","tmNode","parentKey","props","scrollable"]))});return s(),w("div",{class:N([`${o}-dropdown-menu`,t&&`${o}-dropdown-menu--scrollable`]),ref:"bodyRef"},[t?(s(),S(Je,{key:0,contentClass:`${o}-dropdown-menu__content`},{default:()=>n},1032,["contentClass"])):(s(),w(F,{key:1},[P(()=>n)],64)),this.showArrow?(s(),w(F,{key:2},[P(()=>to({clsPrefix:o,arrowStyle:this.arrowStyle,arrowClass:void 0,arrowWrapperClass:void 0,arrowWrapperStyle:void 0}))],64)):P(()=>null)],2)}}),_o=k("dropdown-menu",`
 transform-origin: var(--v-transform-origin);
 background-color: var(--n-color);
 border-radius: var(--n-border-radius);
 box-shadow: var(--n-box-shadow);
 position: relative;
 transition:
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
`,[Ye(),k("dropdown-option",`
 position: relative;
 `,[A("a",`
 text-decoration: none;
 color: inherit;
 outline: none;
 `,[A("&::before",`
 content: "";
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),k("dropdown-option-body",`
 display: flex;
 cursor: pointer;
 position: relative;
 height: var(--n-option-height);
 line-height: var(--n-option-height);
 font-size: var(--n-font-size);
 color: var(--n-option-text-color);
 transition: color .3s var(--n-bezier);
 `,[A("&::before",`
 content: "";
 position: absolute;
 top: 0;
 bottom: 0;
 left: 4px;
 right: 4px;
 transition: background-color .3s var(--n-bezier);
 border-radius: var(--n-border-radius);
 `),ge("disabled",[R("pending",`
 color: var(--n-option-text-color-hover);
 `,[B("prefix, suffix",`
 color: var(--n-option-text-color-hover);
 `),A("&::before","background-color: var(--n-option-color-hover);")]),R("active",`
 color: var(--n-option-text-color-active);
 `,[B("prefix, suffix",`
 color: var(--n-option-text-color-active);
 `),A("&::before","background-color: var(--n-option-color-active);")]),R("child-active",`
 color: var(--n-option-text-color-child-active);
 `,[B("prefix, suffix",`
 color: var(--n-option-text-color-child-active);
 `)])]),R("disabled",`
 cursor: not-allowed;
 opacity: var(--n-option-opacity-disabled);
 `),R("group",`
 font-size: calc(var(--n-font-size) - 1px);
 color: var(--n-group-header-text-color);
 `,[B("prefix",`
 width: calc(var(--n-option-prefix-width) / 2);
 `,[R("show-icon",`
 width: calc(var(--n-option-icon-prefix-width) / 2);
 `)])]),B("prefix",`
 width: var(--n-option-prefix-width);
 display: flex;
 justify-content: center;
 align-items: center;
 color: var(--n-prefix-color);
 transition: color .3s var(--n-bezier);
 z-index: 1;
 `,[R("show-icon",`
 width: var(--n-option-icon-prefix-width);
 `),k("icon",`
 font-size: var(--n-option-icon-size);
 `)]),B("label",`
 white-space: nowrap;
 flex: 1;
 z-index: 1;
 `),B("suffix",`
 box-sizing: border-box;
 flex-grow: 0;
 flex-shrink: 0;
 display: flex;
 justify-content: flex-end;
 align-items: center;
 min-width: var(--n-option-suffix-width);
 padding: 0 8px;
 transition: color .3s var(--n-bezier);
 color: var(--n-suffix-color);
 z-index: 1;
 `,[R("has-submenu",`
 width: var(--n-option-icon-suffix-width);
 `),k("icon",`
 font-size: var(--n-option-icon-size);
 `)]),k("dropdown-menu","pointer-events: all;")]),k("dropdown-offset-container",`
 pointer-events: none;
 position: absolute;
 left: 0;
 right: 0;
 top: -4px;
 bottom: -4px;
 `)]),k("dropdown-divider",`
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-divider-color);
 height: 1px;
 margin: 4px 0;
 `),k("dropdown-menu-wrapper",`
 transform-origin: var(--v-transform-origin);
 width: fit-content;
 `),A(">",[k("scrollbar",`
 height: inherit;
 max-height: inherit;
 `)]),ge("scrollable",`
 padding: var(--n-padding);
 `),R("scrollable",[B("content",`
 padding: var(--n-padding);
 `)])]);const zo={animated:{type:Boolean,default:!0},keyboard:{type:Boolean,default:!0},size:String,inverted:Boolean,placement:{type:String,default:"bottom"},onSelect:[Function,Array],options:{type:Array,default:()=>[]},menuProps:Function,showArrow:Boolean,renderLabel:Function,renderIcon:Function,renderOption:Function,nodeProps:Function,labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},value:[String,Number]},Ko=Object.keys(ye),To={...ye,...zo,...Q.props};var Fo=_({name:"Dropdown",inheritAttrs:!1,props:To,setup(e){const o=O(!1),t=ao(z(e,"show"),o),n=m(()=>{const{keyField:p,childrenField:h}=e;return ro(e.options,{getKey(b){return b[p]},getDisabled(b){return b.disabled===!0},getIgnored(b){return b.type==="divider"||b.type==="render"},getChildren(b){return b[h]}})}),r=m(()=>n.value.treeNodes),a=O(null),d=O(null),c=O(null),l=m(()=>a.value??d.value??c.value??null),y=m(()=>n.value.getPath(l.value).keyPath),f=m(()=>n.value.getPath(e.value).keyPath),C=te(()=>e.keyboard&&t.value);lo({keydown:{ArrowUp:{prevent:!0,handler:le},ArrowRight:{prevent:!0,handler:de},ArrowDown:{prevent:!0,handler:se},ArrowLeft:{prevent:!0,handler:ae},Enter:{prevent:!0,handler:ce},Escape:X}},C);const{mergedClsPrefixRef:g,inlineThemeDisabled:K,mergedComponentPropsRef:T}=me(e),I=m(()=>e.size||T?.value?.Dropdown?.size||"medium"),$=Q("Dropdown","-dropdown",_o,fo,e,g);G(ie,{labelFieldRef:z(e,"labelField"),childrenFieldRef:z(e,"childrenField"),renderLabelRef:z(e,"renderLabel"),renderIconRef:z(e,"renderIcon"),hoverKeyRef:a,keyboardKeyRef:d,lastToggledSubmenuKeyRef:c,pendingKeyPathRef:y,activeKeyPathRef:f,animatedRef:z(e,"animated"),mergedShowRef:t,nodePropsRef:z(e,"nodeProps"),renderOptionRef:z(e,"renderOption"),menuPropsRef:z(e,"menuProps"),doSelect:D,doUpdateShow:j}),he(t,p=>{!e.animated&&!p&&Z()});function D(p,h){const{onSelect:b}=e;b&&pe(b,p,h)}function j(p){const{"onUpdate:show":h,onUpdateShow:b}=e;h&&pe(h,p),b&&pe(b,p),o.value=p}function Z(){a.value=null,d.value=null,c.value=null}function X(){j(!1)}function ae(){U("left")}function de(){U("right")}function le(){U("up")}function se(){U("down")}function ce(){const p=V();p?.isLeaf&&t.value&&(D(p.key,p.rawNode),j(!1))}function V(){const{value:p}=n,{value:h}=l;return!p||h===null?null:p.getNode(h)??null}function U(p){const{value:h}=l,{value:{getFirstAvailableNode:b}}=n;let i=null;if(h===null){const u=b();u!==null&&(i=u.key)}else{const u=V();if(u){let x;switch(p){case"down":x=u.getNext();break;case"up":x=u.getPrev();break;case"right":x=u.getChild();break;case"left":x=u.getParent()}x&&(i=x.key)}}i!==null&&(a.value=null,d.value=i)}const ee=m(()=>{const{inverted:p}=e,h=I.value,{common:{cubicBezierEaseInOut:b},self:i}=$.value,{padding:u,dividerColor:x,borderRadius:H,optionOpacityDisabled:ue,[M("optionIconSuffixWidth",h)]:L,[M("optionSuffixWidth",h)]:De,[M("optionIconPrefixWidth",h)]:Be,[M("optionPrefixWidth",h)]:Ae,[M("fontSize",h)]:Oe,[M("optionHeight",h)]:Fe,[M("optionIconSize",h)]:He}=i,v={"--n-bezier":b,"--n-font-size":Oe,"--n-padding":u,"--n-border-radius":H,"--n-option-height":Fe,"--n-option-prefix-width":Ae,"--n-option-icon-prefix-width":Be,"--n-option-suffix-width":De,"--n-option-icon-suffix-width":L,"--n-option-icon-size":He,"--n-divider-color":x,"--n-option-opacity-disabled":ue};return p?(v["--n-color"]=i.colorInverted,v["--n-option-color-hover"]=i.optionColorHoverInverted,v["--n-option-color-active"]=i.optionColorActiveInverted,v["--n-option-text-color"]=i.optionTextColorInverted,v["--n-option-text-color-hover"]=i.optionTextColorHoverInverted,v["--n-option-text-color-active"]=i.optionTextColorActiveInverted,v["--n-option-text-color-child-active"]=i.optionTextColorChildActiveInverted,v["--n-prefix-color"]=i.prefixColorInverted,v["--n-suffix-color"]=i.suffixColorInverted,v["--n-group-header-text-color"]=i.groupHeaderTextColorInverted):(v["--n-color"]=i.color,v["--n-option-color-hover"]=i.optionColorHover,v["--n-option-color-active"]=i.optionColorActive,v["--n-option-text-color"]=i.optionTextColor,v["--n-option-text-color-hover"]=i.optionTextColorHover,v["--n-option-text-color-active"]=i.optionTextColorActive,v["--n-option-text-color-child-active"]=i.optionTextColorChildActive,v["--n-prefix-color"]=i.prefixColor,v["--n-suffix-color"]=i.suffixColor,v["--n-group-header-text-color"]=i.groupHeaderTextColor),v}),Y=K?ke("dropdown",m(()=>`${I.value[0]}${e.inverted?"i":""}`),ee,e):void 0;return{mergedClsPrefix:g,mergedTheme:$,mergedSize:I,tmNodes:r,mergedShow:t,handleAfterLeave:()=>{e.animated&&Z()},doUpdateShow:j,cssVars:K?void 0:ee,themeClass:Y?.themeClass,onRender:Y?.onRender}},render(){const e=(n,r,a,d,c)=>{const{mergedClsPrefix:l,menuProps:y}=this;this.onRender?.();const f=y?.(void 0,this.tmNodes.map(g=>g.rawNode))||{},C={ref:co(r),class:[n,`${l}-dropdown`,`${l}-dropdown--${this.mergedSize}-size`,this.themeClass],clsPrefix:l,tmNodes:this.tmNodes,style:[...a,this.cssVars],showArrow:this.showArrow,arrowStyle:this.arrowStyle,scrollable:this.scrollable,onMouseenter:d,onMouseleave:c};return E($e,J(this.$attrs,C,f))},{mergedTheme:o}=this,t={show:this.mergedShow,theme:o.peers.Popover,themeOverrides:o.peerOverrides.Popover,internalOnAfterLeave:this.handleAfterLeave,internalRenderBody:e,onUpdateShow:this.doUpdateShow,"onUpdate:show":void 0};return s(),S(_e,Ze(this.$props,Ko,t),{_:1,trigger:Re(()=>this.$slots.default?.())},16)}});export{so as C,Fo as D,So as N,Oo as T,co as c,fo as d,mo as t};
