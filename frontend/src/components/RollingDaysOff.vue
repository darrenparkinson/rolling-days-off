<template>
  <v-container>
    <v-row>
      <v-col cols="12" sm="6" md="4">
        <v-row>
          <v-col>
            <v-menu
              ref="menu"
              v-model="menu"
              :close-on-content-click="false"
              :return-value.sync="startDate"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="startDate"
                  label="Most recent day off"
                  prepend-icon="mdi-calendar"
                  readonly
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker v-model="startDate" scrollable>
                <v-spacer></v-spacer>
                <v-btn text color="primary" @click="menu = false">Cancel</v-btn>
                <v-btn text color="primary" @click="$refs.menu.save(startDate)">OK</v-btn>
              </v-date-picker>
            </v-menu>
          </v-col>
        </v-row>
        <v-row>
          <!-- <v-spacer></v-spacer> -->

          <v-col>
            <v-menu
              ref="menu2"
              v-model="menu2"
              :close-on-content-click="false"
              :return-value.sync="endDate"
              transition="scale-transition"
              offset-y
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="endDate"
                  label="Find dates up to"
                  prepend-icon="mdi-calendar"
                  readonly
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker v-model="endDate" scrollable>
                <v-spacer></v-spacer>
                <v-btn text color="primary" @click="menu2 = false">Cancel</v-btn>
                <v-btn text color="primary" @click="$refs.menu2.save(endDate)">OK</v-btn>
              </v-date-picker>
            </v-menu>
          </v-col>
        </v-row>
        <v-row>
          <!-- <v-spacer></v-spacer> -->
          <v-col>
            <v-radio-group v-model="dayOff" :mandatory="true" column>
              <v-text-field
                v-model="dayOff"
                label="Persistent Day Off"
                prepend-icon="mdi-calendar-remove"
                readonly
              ></v-text-field>
              <v-radio label="Monday" value="Monday"></v-radio>
              <v-radio label="Tuesday" value="Tuesday"></v-radio>
              <v-radio label="Wednesday" value="Wednesday"></v-radio>
              <v-radio label="Thursday" value="Thursday"></v-radio>
              <v-radio label="Friday" value="Friday"></v-radio>
              <v-radio label="Saturday" value="Saturday"></v-radio>
              <v-radio label="Sunday" value="Sunday"></v-radio>
            </v-radio-group>
          </v-col>
        </v-row>
        <v-row>
          <v-btn v-on:click="showDates" color="primary">Show Dates</v-btn>
        </v-row>
      </v-col>
      <v-col>
        <days-list :daysOff="daysOff" :loading="loading" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import DaysList from "@/components/DaysList.vue";
import axios from "axios";
export default {
  components: {
    DaysList
  },
  data: () => ({
    startDate: new Date().toISOString().substr(0, 10),
    endDate: new Date(new Date().getFullYear() + 1, 3, 1)
      .toISOString()
      .substr(0, 10),
    dayOff: "Sunday",
    menu: false,
    modal: false,
    menu2: false,
    daysOff: [],
    loading: false
  }),
  methods: {
    async showDates() {
      console.log("Showing Dates");
      console.log(this.startDate);
      console.log(this.endDate);
      console.log(this.dayOff);
      let dayOffNum = 0;
      switch (this.dayOff) {
        case "Sunday":
          dayOffNum = 0;
          break;
        case "Monday":
          dayOffNum = 1;
          break;
        case "Tuesday":
          dayOffNum = 2;
          break;
        case "Wednesday":
          dayOffNum = 3;
          break;
        case "Thursday":
          dayOffNum = 4;
          break;
        case "Friday":
          dayOffNum = 5;
          break;
        case "Saturday":
          dayOffNum = 6;
          break;
        default:
          break;
      }
      console.log(dayOffNum);
      this.loading = true;
      try {
        const response = await axios.get(
          `/api/v1/daysoff?startDate=${this.startDate}&endDate=${this.endDate}&dayOff=${dayOffNum}`
        );
        this.daysOff = response.data;
        this.loading = false;
      } catch (error) {
        console.log(error);
        this.loading = false;
      }
    }
  }
};
</script>

<style>
/* Fix spacing between dates and persistent day off */
.v-input--selection-controls {
  margin-top: 0;
}
</style>