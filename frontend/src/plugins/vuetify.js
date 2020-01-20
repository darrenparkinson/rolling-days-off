import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
            light: {
                primary: '#c70f15',
                secondary: '#DC3232',
                accent: '#DC3232',
                error: '#b71c1c',
            },
        },
    },
});
