import{A as ve,U as ge,D as J,av as a,aw as X,C as U,at as s,_ as Y,d as me,E as Q,bJ as A,o as b,c as y,a as k,G as l,b9 as w,ac as o,a9 as q,J as pe,az as we,k as N,a5 as xe,x as M,a7 as j,n as G,b4 as ye,F as ke,aq as Se,L as x,M as E,b2 as d,ae as Ce}from"./index-C8SHz2c_.js";import{b as Be}from"./FormItem-BA7pYDWv.js";var _e={buttonHeightSmall:"14px",buttonHeightMedium:"18px",buttonHeightLarge:"22px",buttonWidthSmall:"14px",buttonWidthMedium:"18px",buttonWidthLarge:"22px",buttonWidthPressedSmall:"20px",buttonWidthPressedMedium:"24px",buttonWidthPressedLarge:"28px",railHeightSmall:"18px",railHeightMedium:"22px",railHeightLarge:"26px",railWidthSmall:"32px",railWidthMedium:"40px",railWidthLarge:"48px"};function Re(e){const{primaryColor:c,opacityDisabled:f,borderRadius:n,textColor3:S}=e;return{..._e,iconColor:S,textColor:"white",loadingColor:c,opacityDisabled:f,railColor:"rgba(0, 0, 0, .14)",railColorActive:c,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:n,railBorderRadiusMedium:n,railBorderRadiusLarge:n,buttonBorderRadiusSmall:n,buttonBorderRadiusMedium:n,buttonBorderRadiusLarge:n,boxShadowFocus:`0 0 0 2px ${ge(c,{alpha:.2})}`}}const ze={common:ve,self:Re};var $e=J("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[a("children-placeholder",`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),a("rail-placeholder",`
 display: flex;
 flex-wrap: none;
 `),a("button-placeholder",`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),J("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[X({left:"50%",top:"50%",originalTransform:"translateX(-50%) translateY(-50%)"})]),a("checked, unchecked",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 box-sizing: border-box;
 position: absolute;
 white-space: nowrap;
 top: 0;
 bottom: 0;
 display: flex;
 align-items: center;
 line-height: 1;
 `),a("checked",`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),a("unchecked",`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),U("&:focus",[a("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),s("round",[a("rail","border-radius: calc(var(--n-rail-height) / 2);",[a("button","border-radius: calc(var(--n-button-height) / 2);")])]),Y("disabled",[Y("icon",[s("rubber-band",[s("pressed",[a("rail",[a("button","max-width: var(--n-button-width-pressed);")])]),a("rail",[U("&:active",[a("button","max-width: var(--n-button-width-pressed);")])]),s("active",[s("pressed",[a("rail",[a("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])]),a("rail",[U("&:active",[a("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])])])])])]),s("active",[a("rail",[a("button","left: calc(100% - var(--n-button-width) - var(--n-offset))")])]),a("rail",`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[a("button-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 font-size: calc(var(--n-button-height) - 4px);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 display: flex;
 justify-content: center;
 align-items: center;
 line-height: 1;
 `,[X()]),a("button",`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-height);
 width: var(--n-button-width-pressed);
 max-width: var(--n-button-width);
 border-radius: var(--n-button-border-radius);
 background-color: var(--n-button-color);
 box-shadow: var(--n-button-box-shadow);
 box-sizing: border-box;
 cursor: inherit;
 content: "";
 position: absolute;
 transition:
 background-color .3s var(--n-bezier),
 left .3s var(--n-bezier),
 opacity .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 `)]),s("active",[a("rail","background-color: var(--n-rail-color-active);")]),s("loading",[a("rail",`
 cursor: wait;
 `)]),s("disabled",[a("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)])]);const Ve=["aria-checked","tabindex","onClick","onFocus","onBlur","onKeyup","onKeydown"],Fe={...Q.props,size:String,value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},spinProps:Object,onChange:[Function,Array]};let z;var Me=me({name:"Switch",props:Fe,slots:Object,setup(e){z===void 0&&(typeof CSS<"u"?typeof CSS.supports<"u"?z=CSS.supports("width","max(1px)"):z=!1:z=!0);const{mergedClsPrefixRef:c,inlineThemeDisabled:f,mergedComponentPropsRef:n}=pe(e),S=Q("Switch","-switch",$e,ze,e,c),v=we(e,{mergedSize(t){if(e.size!==void 0)return e.size;if(t)return t.mergedSize.value;const p=n?.value?.Switch?.size;return p||"medium"}}),{mergedSizeRef:C,mergedDisabledRef:g}=v,B=N(e.defaultValue),$=Ce(e,"value"),m=Be($,B),V=M(()=>m.value===e.checkedValue),i=N(!1),r=N(!1),_=M(()=>{const{railStyle:t}=e;if(t)return t({focused:r.value,checked:V.value})});function F(t){const{"onUpdate:value":p,onChange:W,onUpdateValue:P}=e,{nTriggerFormInput:H,nTriggerFormChange:L}=v;p&&j(p,t),P&&j(P,t),W&&j(W,t),B.value=t,H(),L()}function Z(){const{nTriggerFormFocus:t}=v;t()}function ee(){const{nTriggerFormBlur:t}=v;t()}function te(){e.loading||g.value||(m.value!==e.checkedValue?F(e.checkedValue):F(e.uncheckedValue))}function ae(){r.value=!0,Z()}function ie(){r.value=!1,ee(),i.value=!1}function oe(t){e.loading||g.value||t.key===" "&&(m.value!==e.checkedValue?F(e.checkedValue):F(e.uncheckedValue),i.value=!1)}function ne(t){e.loading||g.value||t.key===" "&&(t.preventDefault(),i.value=!0)}const O=M(()=>{const{value:t}=C,{self:{opacityDisabled:p,railColor:W,railColorActive:P,buttonBoxShadow:H,buttonColor:L,boxShadowFocus:re,loadingColor:le,textColor:se,iconColor:de,[x("buttonHeight",t)]:u,[x("buttonWidth",t)]:ce,[x("buttonWidthPressed",t)]:ue,[x("railHeight",t)]:h,[x("railWidth",t)]:R,[x("railBorderRadius",t)]:he,[x("buttonBorderRadius",t)]:be},common:{cubicBezierEaseInOut:fe}}=S.value;let T,D,K;return z?(T=`calc((${h} - ${u}) / 2)`,D=`max(${h}, ${u})`,K=`max(${R}, calc(${R} + ${u} - ${h}))`):(T=E((d(h)-d(u))/2),D=E(Math.max(d(h),d(u))),K=d(h)>d(u)?R:E(d(R)+d(u)-d(h))),{"--n-bezier":fe,"--n-button-border-radius":be,"--n-button-box-shadow":H,"--n-button-color":L,"--n-button-width":ce,"--n-button-width-pressed":ue,"--n-button-height":u,"--n-height":D,"--n-offset":T,"--n-opacity-disabled":p,"--n-rail-border-radius":he,"--n-rail-color":W,"--n-rail-color-active":P,"--n-rail-height":h,"--n-rail-width":R,"--n-width":K,"--n-box-shadow-focus":re,"--n-loading-color":le,"--n-text-color":se,"--n-icon-color":de}}),I=f?xe("switch",M(()=>C.value[0]),O,e):void 0;return{handleClick:te,handleBlur:ie,handleFocus:ae,handleKeyup:oe,handleKeydown:ne,mergedRailStyle:_,pressed:i,mergedClsPrefix:c,mergedValue:m,checked:V,mergedDisabled:g,cssVars:f?void 0:O,themeClass:I?.themeClass,onRender:I?.onRender}},render(){const{mergedClsPrefix:e,mergedDisabled:c,checked:f,mergedRailStyle:n,onRender:S,$slots:v}=this;S?.();const{checked:C,unchecked:g,icon:B,"checked-icon":$,"unchecked-icon":m}=v,V=!(A(B)&&A($)&&A(m));return b(),y("div",{role:"switch","aria-checked":f,class:o([`${e}-switch`,this.themeClass,V&&`${e}-switch--icon`,f&&`${e}-switch--active`,c&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`]),tabindex:this.mergedDisabled?void 0:0,style:q(this.cssVars),onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},[k("div",{class:o(`${e}-switch__rail`),"aria-hidden":"true",style:q(n)},[l(()=>w(C,i=>w(g,r=>i||r?(b(),y("div",{key:4,"aria-hidden":!0,class:o(`${e}-switch__children-placeholder`)},[k("div",{class:o(`${e}-switch__rail-placeholder`)},[k("div",{class:o(`${e}-switch__button-placeholder`)},null,2),l(()=>i)],2),k("div",{class:o(`${e}-switch__rail-placeholder`)},[k("div",{class:o(`${e}-switch__button-placeholder`)},null,2),l(()=>r)],2)],2)):null))),k("div",{class:o(`${e}-switch__button`)},[l(()=>w(B,i=>w($,r=>w(m,_=>(b(),G(Se,null,{default:()=>this.loading?(b(),G(ye,ke({key:"loading",clsPrefix:e,strokeWidth:20},this.spinProps),null,16,["clsPrefix"])):this.checked&&(r||i)?(b(),y("div",{class:o(`${e}-switch__button-icon`),key:r?"checked-icon":"icon"},[l(()=>r||i)],2)):!this.checked&&(_||i)?(b(),y("div",{class:o(`${e}-switch__button-icon`),key:_?"unchecked-icon":"icon"},[l(()=>_||i)],2)):null},1024)))))),l(()=>w(C,i=>i&&(b(),y("div",{key:"checked",class:o(`${e}-switch__checked`)},[l(()=>i)],2)))),l(()=>w(g,i=>i&&(b(),y("div",{key:"unchecked",class:o(`${e}-switch__unchecked`)},[l(()=>i)],2))))],2)],6)],46,Ve)}});export{Me as S};
