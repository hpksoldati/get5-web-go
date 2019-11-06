import Vue from 'vue'
import Router from 'vue-router'
import axios from 'axios';
import VueAxios from 'vue-axios'

import Matches from '@/components/Matches'
import Match from '@/components/Match'
import MyMatches from '@/components/MyMatches'
import Metrics from '@/components/Metrics'
import Team from '@/components/Team'

Vue.use(VueAxios, axios)
Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: '/matches'
    },
    {
      path: '/matches',
      name: 'Matches',
      component: Matches
    },
    {
      path: '/mymatches',
      name: 'My matches',
      component: MyMatches
    },
    {
      path: '/match',
      name: 'Match',
      component: Match
    },
    {
      path: '/metrics',
      name: 'Metrics',
      component: Metrics
    },
    {
      path: '/team',
      name: 'Team',
      component: Team
    }
  ]
})
