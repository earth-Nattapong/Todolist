<template>
  <div class="container">
    <div class="panel">
      <div class="panel-heading">
        <div class="columns">
          <div class="column is-four-fifths mt-2 ml-2">
            <p class="has-text-weight-semibold is-size-3 ">
              ToDoList 
            </p>
            <p class="has-text-weight-semibold is-size-6 mt-3 ml-1 ">
              To Day : {{ time }}
            </p>
          </div>
          <div class="column">
            <button
              class="button is-primary"
              id="push"
              @click="checkmodal = true"
            >
              Add
            </button>
          </div>
        </div>
        <div class="box">
          <div class="columns">
            <div class="column is-6">Completed : {{ numComplete }} <i class="fa-solid fa-square-check has-text-primary" ></i></div>
            <div class="column is-6">Incomplete : {{ numIncomplete }} <i class="fa-solid fa-circle-xmark has-text-danger"></i></div>
          </div>
        </div>
      </div>
      <a
        v-for="task in taks"
        :key="task.id"
        class="panel-block is-active"
        @click="editTask(task)"
        :style="{
          'text-decoration': task.check === 'susess' ? 'line-through' : 'none',
          color: task.check === 'susess' ? 'green' : 'initial',
        }"
      >
        <span class="panel-icon">
          <i class="fas fa-edit"></i>
        </span>
        {{ task.taskname }}
        <span class="icon" key="true">
          <i
            v-if="task.importance == 1"
            class="fas fa-check-circle has-text-primary"
          ></i>
          <i
            v-if="task.importance == 2"
            class="fas fa-check-circle has-text-warning"
          ></i>
          <i
            v-if="task.importance == 3"
            class="fas fa-check-circle has-text-danger"
          ></i>
        </span>
      </a>
    </div>
    <div class="modal" :class="{ 'is-active': checkmodal }">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">New Task</p>
        </header>
        <section class="modal-card-body">
          <p class="label">Task Name</p>
          <input class="input is-rounded" type="text" v-model="addtaskname" />
          <p class="label">Title</p>
          <input class="input is-rounded" type="text" v-model="addtitle" />
          <p class="label">Importence Score</p>
          <div class="select is-rounded">
            <select
              class="is-hovered"
              name="score"
              id="score"
              v-model="scoretask"
            >
              <option value="1">1</option>
              <option value="2">2</option>
              <option value="3">3</option>
            </select>
          </div>
        </section>
        <footer class="modal-card-foot">
          <button class="button is-success" @click="addTask()">Save</button>
          <button class="button is-danger is-light" @click="checkmodal = false">
            Cancel
          </button>
        </footer>
      </div>
    </div>
    <div class="modal" :class="{ 'is-active': checkeditmodal }">
      <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">{{ tasknameedit }}</p>
          <div v-if="checksusess == true">
            <button
              class="button is-success"
              @click="updatesusess(this.editID)"
            >
              Susess
            </button>
          </div>
        </header>
        <section class="modal-card-body">
          <p class="label">Title</p>
          <div>
            <p class="text" v-if="checkeditbutton == false">
              {{ titleedit }}
            </p>
            <input
              v-if="checkeditbutton == true"
              class="input is-rounded"
              type="text"
              v-model="this.titleedit"
            />
          </div>
          <div class="select is-rounded mt-3 is-small">
            <select
              class="is-hovered"
              name="score"
              id="score"
              v-model="scoretask"
              
            >
              <option value="1">1</option>
              <option value="2">2</option>
              <option value="3">3</option>
            </select>
          </div>
        </section>
        <footer class="modal-card-foot">
          <button
            class="button is-info is-light"
            @click="checkeditbutton = true"
            v-if="checkeditbutton == false"
          >
            Edit
          </button>
          <button
            class="button is-success is-light"
            @click="saveEditTask(this.editID)"
            v-if="checkeditbutton == true"
          >
            Save
          </button>
          <button
            class="button is-danger is-light"
            @click="(checkeditmodal = false), (checkeditbutton = false)"
          >
            Cancel
          </button>
          <button class="button is-danger m-1" @click="deleteTask(this.editID)">
            Delete<i class="fas fa-trash-alt"></i>
          </button>
        </footer>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { useNow, useDateFormat } from "@vueuse/core";

const formatted = useDateFormat(useNow(), "YYYY/MM/DD HH:mm:ss");
export default {
  name: "ToDoList",
  data() {
    return {
      taks: [],
      time: formatted,
      addtaskname: "",
      checkmodal: false,
      addtitle: "",
      scoretask: 0,
      //edit
      checkeditmodal: false,
      tasknameedit: "",
      titleedit: "",
      checkeditbutton: false,
      editID: 0,
      checksusess: false,
    };
  },
  mounted() {
    this.getTask();
  },
  computed: {
    numComplete() {
      const count = this.taks.filter((val) => val.check == "susess");
      return count.length;
    },
    numIncomplete() {
      return this.taks.length - this.numComplete;
    },
    filteredTask() {
      function comparealpa(a, b) {
        if (a.importance > b.importance) {
          return -1;
        } else if (a.importance < b.importance) {
          return 1;
        } else {
          return 0;
        }
      }
      this.taks.sort(comparealpa);
    },
  },
  methods: {
    getTask() {
      axios
        .get("http://localhost:8080/tasks")
        .then((response) => {
          this.taks = response.data;
          console.log(this.taks);
          this.filteredTask();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    addTask() {
      if (this.addtaskname != "" && this.addtitle != "") {
        const data = {
          taskname: this.addtaskname,
          title: this.addtitle,
          check: "unsusess",
          importance: parseInt(this.scoretask),
        };
        axios
          .post("http://localhost:8080/tasks", data)
          .then((response) => {
            this.getTask();
          })
          .catch((err) => {
            console.log(err);
          });
        this.checkmodal = false;
        this.addtaskname = "";
        this.addtitle = "";
        this.scoretask = 0;
      } else {
        alert("กรุณากรอกข้อมูลให้ครบถ้วน");
      }
    },
    editTask(value) {
      this.checkeditmodal = true;
      this.tasknameedit = value.taskname;
      this.titleedit = value.title;
      this.editID = value.id;
      this.scoretask = value.importance;
      if (value.check == "susess") {
        this.checksusess = false;
      } else {
        this.checksusess = true;
      }
    },
    saveEditTask(taskId) {
      if (this.tasknameedit != "" && this.titleedit != "") {
        const data = {
          taskname: this.tasknameedit,
          title: this.titleedit,
          check: "unsusess",
          importance: parseInt(this.scoretask),
        };
        console.log(data);
        axios
          .put(`http://localhost:8080/tasks/${taskId}`, data)
          .then((response) => {
            //   this.taks = response.data;
            //   console.log(this.taks);
            this.getTask();
          })
          .catch((err) => {
            console.log(err);
          });
        this.checkeditmodal = false;
        this.tasknameedit = "";
        this.titleedit = "";
        this.checkeditbutton = false;
        this.editID = 0;
      } else {
        alert("กรุณากรอกข้อมูลให้ครบถ้วน");
      }
    },
    deleteTask(taskId) {
      axios
        .delete(`http://localhost:8080/tasks/${taskId}`)
        .then((response) => {
          //   this.taks = response.data;
          //   console.log(this.taks);
          this.getTask();
        })
        .catch((err) => {
          console.log(err);
        });
      this.checkeditmodal = false;
      this.tasknameedit = "";
      this.titleedit = "";
      this.checkeditbutton = false;
      this.editID = 0;
    },
    updatesusess(taskId) {
      const data = {
        taskname: this.tasknameedit,
        title: this.titleedit,
        check: "susess",
        importance: this.scoretask,
      };
      axios
        .put(`http://localhost:8080/tasks/${taskId}`, data)
        .then((response) => {
          //   this.taks = response.data;
          //   console.log(this.taks);
          this.getTask();
        })
        .catch((err) => {
          console.log(err);
        });
      this.checkeditmodal = false;
      this.tasknameedit = "";
      this.titleedit = "";
      this.checkeditbutton = false;
      this.editID = 0;
    },
  },
};
</script>

<style>
@import "https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css";
@import url("https://fonts.googleapis.com/css2?family=Roboto:wght@400;500&display=swap");

body {
  background: #75cb8f;
  font-family: "Roboto", sans-serif;
}
.container {
  width: 200%;
  top: 5%;
  background: white;
  border-radius: 20px;
}
</style>