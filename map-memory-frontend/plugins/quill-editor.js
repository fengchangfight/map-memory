import { VueEditor,Quill } from "vue2-editor";
import Vue from 'vue';
Vue.component('VueEditor', VueEditor)

const Block = Quill.import('blots/block');
Block.tagName = 'DIV';
Quill.register(Block, true);
