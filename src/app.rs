use std::error;

use cursorvec::CursorVec;
use tui_tree_widget::TreeItem;

use crate::items::{Item, StatefulTree};

/// Application result type.
pub type AppResult<T> = std::result::Result<T, Box<dyn error::Error>>;

/// Application.
#[derive(Debug)]
pub struct App {
    pub running: bool,
    pub counter: u8,
    pub sidebar: SideBar,
    pub settings: Settings,
}

#[derive(Debug)]
pub struct SideBar {
    pub size: u16,
    pub selected: usize,
    pub tree: StatefulTree<'static>,
}

#[derive(Debug)]
pub struct Settings {
    pub show_sidebar: bool,
    pub show_help: bool,
}

impl Default for App {
    fn default() -> Self {
        let mut tree = StatefulTree::with_items(vec![
            TreeItem::new_leaf(Item::new("a")),
            TreeItem::new(
                Item::new("b"),
                vec![
                    TreeItem::new_leaf(Item::new("c")),
                    TreeItem::new(
                        Item::new("d"),
                        vec![
                            TreeItem::new_leaf(Item::new("e")),
                            TreeItem::new_leaf(Item::new("f")),
                        ],
                    ),
                    TreeItem::new_leaf(Item::new("g")),
                ],
            ),
            TreeItem::new_leaf(Item::new("d")),
        ]);
        tree.first();
        Self {
            running: true,
            counter: 0,
            sidebar: SideBar {
                size: 45,
                selected: 0,
                tree,
            },
            settings: Settings {
                show_sidebar: true,
                show_help: false,
            },
        }
    }
}

impl App {
    /// Constructs a new instance of [`App`].
    pub fn new() -> Self {
        Self::default()
    }

    /// Handles the tick event of the terminal.
    pub fn tick(&self) {}

    /// Set running to false to quit the application.
    pub fn quit(&mut self) {
        self.running = false;
    }

    pub fn toggle_sidebar(&mut self) {
        self.settings.show_sidebar = !self.settings.show_sidebar;
    }

    pub fn sidebar_size(&self) -> u16 {
        match self.settings.show_sidebar {
            true => self.sidebar.size,
            false => 0,
        }
    }
}
