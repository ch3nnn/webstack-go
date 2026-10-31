import{n as w,M as _,v as l,as as T,x as y,d as $,bZ as B,y as S,o as f,c as h,C as m,A as r,a as z,z as v,D as R,b1 as V,b as N,G as P,I as D,a1 as I,J as g,k as L,aB as O,ac as W,aj as j,aO as b}from"./index-BwLaMb6L.js";import{u as E}from"./use-compitable-D4GrHVKv.js";function H(e){const{opacityDisabled:t,heightTiny:s,heightSmall:d,heightMedium:n,heightLarge:o,heightHuge:p,primaryColor:c,fontSize:i}=e;return{fontSize:i,textColor:c,sizeTiny:s,sizeSmall:d,sizeMedium:n,sizeLarge:o,sizeHuge:p,color:c,opacitySpinning:t}}const M={common:w,self:H};var A=_([_("@keyframes spin-rotate",`
 from {
 transform: rotate(0);
 }
 to {
 transform: rotate(360deg);
 }
 `),l("spin-container",`
 position: relative;
 `,[l("spin-body",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[T()])]),l("spin-body",`
 display: inline-flex;
 align-items: center;
 justify-content: center;
 flex-direction: column;
 `),l("spin",`
 display: inline-flex;
 height: var(--n-size);
 width: var(--n-size);
 font-size: var(--n-size);
 color: var(--n-color);
 `,[y("rotate",`
 animation: spin-rotate 2s linear infinite;
 `)]),l("spin-description",`
 display: inline-block;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 margin-top: 8px;
 `),l("spin-content",`
 opacity: 1;
 transition: opacity .3s var(--n-bezier);
 pointer-events: all;
 `,[y("spinning",`
 user-select: none;
 -webkit-user-select: none;
 pointer-events: none;
 opacity: var(--n-opacity-spinning);
 `)])]);const K={small:20,medium:18,large:16},F={...S.props,contentClass:String,contentStyle:[Object,String],description:String,size:{type:[String,Number],default:"medium"},show:{type:Boolean,default:!0},rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0},delay:Number,...B,strokeWidth:Number};var X=$({name:"Spin",props:F,slots:Object,setup(e){const{mergedClsPrefixRef:t,inlineThemeDisabled:s}=P(e),d=S("Spin","-spin",A,M,e,t),n=g(()=>{const{size:i}=e,{common:{cubicBezierEaseInOut:a},self:u}=d.value,{opacitySpinning:k,color:C,textColor:x}=u;return{"--n-bezier":a,"--n-opacity-spinning":k,"--n-size":typeof i=="number"?W(i):u[j("size",i)],"--n-color":C,"--n-text-color":x}}),o=s?D("spin",g(()=>{const{size:i}=e;return typeof i=="number"?String(i):i[0]}),n,e):void 0,p=E(e,["spinning","show"]),c=L(!1);return I(i=>{let a;if(p.value){const{delay:u}=e;if(u){a=window.setTimeout(()=>{c.value=!0},u),i(()=>{clearTimeout(a)});return}}c.value=p.value}),{mergedClsPrefix:t,active:c,mergedStrokeWidth:g(()=>{const{strokeWidth:i}=e;if(i!==void 0)return i;const{size:a}=e;return K[typeof a=="number"?"medium":a]}),cssVars:s?void 0:n,themeClass:o?.themeClass,onRender:o?.onRender}},render(){const{$slots:e,mergedClsPrefix:t,description:s}=this,d=e.icon&&this.rotate,n=(s||e.description)&&(f(),h("div",{class:r(`${t}-spin-description`)},[m(()=>s||e.description?.())],2)),o=e.icon?(f(),h("div",{key:1,class:r([`${t}-spin-body`,this.themeClass])},[z("div",{class:r([`${t}-spin`,d&&`${t}-spin--rotate`]),style:v(e.default?"":this.cssVars)},[m(()=>e.icon())],6),m(()=>n)],2)):(f(),h("div",{key:2,class:r([`${t}-spin-body`,this.themeClass])},[(f(),R(V,{clsPrefix:t,style:v(e.default?"":this.cssVars),stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,radius:this.radius,scale:this.scale,class:r(`${t}-spin`)},null,8,["clsPrefix","style","stroke","stroke-width","radius","scale","class"])),m(()=>n)],2));return this.onRender?.(),e.default?(f(),h("div",{key:3,class:r([`${t}-spin-container`,this.themeClass]),style:v(this.cssVars)},[z("div",{class:r([`${t}-spin-content`,this.active&&`${t}-spin-content--spinning`,this.contentClass]),style:v(this.contentStyle)},[m(()=>e.default?.())],6),N(O,{name:"fade-in-transition"},{default:()=>this.active?o:null},1024)],6)):o}});async function Y(){const{data:e}=await b.get("/admin/config");return e}async function Z(e){const t=new FormData;e.is_search!==void 0&&t.append("is_search",String(e.is_search)),e.site_title!==void 0&&t.append("site_title",e.site_title),e.site_keyword!==void 0&&t.append("site_keyword",e.site_keyword),e.site_desc!==void 0&&t.append("site_desc",e.site_desc),e.site_record!==void 0&&t.append("site_record",e.site_record),e.site_url!==void 0&&t.append("site_url",e.site_url),e.logo&&t.append("logo",e.logo),e.favicon&&t.append("favicon",e.favicon),await b.put("/admin/config",t)}async function q(e){const t={};e.about_site!==void 0&&(t.about_site=e.about_site),e.about_author!==void 0&&(t.about_author=e.about_author),e.is_about!==void 0&&(t.is_about=e.is_about),await b.put("/admin/config",t)}export{X as S,Z as a,Y as f,q as u};
