#!/usr/bin/env node

import { Screen } from "neo-blessed";

let mainScreen: Screen;

/**
 * Create main application blessed screen and table objects.
 */
export function init() {
  mainScreen = Screen({
    smartCSR: true,
  });

  mainScreen.title = "neoss";
  mainScreen.key(["escape", "q", "C-c"], function () {
    return process.exit(0);
  });

  /*
  table = Table({
    keys: true,
    interactive: true,
    tags: true,
    top: "0",
    left: "center",
    width: "100%",
    height: "shrink",
    border: {
      type: "line",
    },
    style: {
      fg: "white",
      border: {
        fg: "white",
      },
      focus: {
        bg: "blue",
      },
      header: {
        fg: "black",
        bg: "white",
      },
    },
  });
  */
}
