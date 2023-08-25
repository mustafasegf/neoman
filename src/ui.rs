use ratatui::{
    prelude::*,
    widgets::{Block, BorderType, Borders, Paragraph, Tabs, Wrap},
};
use strum::IntoEnumIterator;
use tui_tree_widget::Tree;

use crate::app::{App, RequestMenu};

/// Renders the user interface widgets.
pub fn render<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>) {
    let chunks = Layout::default()
        .direction(Direction::Horizontal)
        .constraints([Constraint::Length(app.sidebar_size()), Constraint::Min(0)].as_ref())
        .split(frame.size());

    // render side bar
    sidebar(app, frame, chunks[0]);
    // render main bar
    mainbar(app, frame, chunks[1]);

    // frame.render_widget(
    //     Paragraph::new(format!(
    //         "This is a tui template.\n\
    //             Press `Esc`, `Ctrl-C` or `q` to stop running.\n\
    //             Press left and right to increment and decrement the counter respectively.\n\
    //             Counter: {}",
    //         app.counter
    //     ))
    //     .block(
    //         Block::default()
    //             .title("Template")
    //             .title_alignment(Alignment::Center)
    //             .borders(Borders::ALL)
    //             .border_type(BorderType::Rounded),
    //     )
    //     .style(Style::default().fg(Color::Cyan).bg(Color::Black))
    //     .alignment(Alignment::Center),
    //     frame.size(),
    // )
}

pub fn sidebar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let indicies = app.sidebar.tree.state.selected();
    let item = app.sidebar.tree.items.get(indicies[0]);

    let selected = indicies.iter().skip(1).fold(item, |item, &i| {
        item.and_then(|item| item.children().get(i))
    });

    let block = Block::default()
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded)
        .title(format!(
            "{:?}",
            selected.map(|item| item.inner().to_string())
        ));

    let items = Tree::new(app.sidebar.tree.items.clone())
        .block(block)
        .highlight_style(
            Style::default()
                .fg(Color::LightBlue)
                .add_modifier(Modifier::BOLD)
                .add_modifier(Modifier::UNDERLINED),
        );
    frame.render_stateful_widget(items, area, &mut app.sidebar.tree.state);
}

pub fn mainbar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let chunks = Layout::default()
        .direction(Direction::Vertical)
        .constraints(
            [
                Constraint::Min(3),
                Constraint::Min(3),
                Constraint::Ratio(1, 2),
                Constraint::Ratio(1, 2),
            ]
            .as_ref(),
        )
        .split(area);

    let block = Block::default()
        .title("Reponse")
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded);

    tabs(app, frame, chunks[0]);
    urlbar(app, frame, chunks[1]);
    requestbar(app, frame, chunks[2]);
    frame.render_widget(block, chunks[3]);
}

pub fn tabs<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let titles = app
        .tabs
        .tabs
        .iter()
        .map(|item| Line::from(item.to_string()))
        .collect();

    let tabs = Tabs::new(titles)
        .block(Block::default().borders(Borders::ALL).title("Tabs"))
        .select(app.tabs.selected)
        .style(Style::default().fg(Color::Cyan))
        .highlight_style(
            Style::default()
                .add_modifier(Modifier::BOLD)
                .bg(Color::Black),
        );

    frame.render_widget(tabs, area);
}

pub fn urlbar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let chunks = Layout::default()
        .direction(Direction::Horizontal)
        .constraints([Constraint::Min(10), Constraint::Min(0)].as_ref())
        .split(area);

    let block = Block::default()
        .title(format!("URL: {}", app.urlbar.title))
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded);

    let text = Paragraph::new(app.urlbar.text.clone())
        .block(block)
        .wrap(Wrap { trim: true })
        .alignment(Alignment::Left);

    let method_block = Block::default()
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded);

    let method = Paragraph::new(app.urlbar.method.to_string())
        .block(method_block)
        .wrap(Wrap { trim: true })
        .alignment(Alignment::Left);

    frame.render_widget(method, chunks[0]);
    frame.render_widget(text, chunks[1]);
}

pub fn requestbar<B: Backend>(app: &mut App, frame: &mut Frame<'_, B>, area: Rect) {
    let chunks = Layout::default()
        .direction(Direction::Vertical)
        .constraints([Constraint::Min(1), Constraint::Min(0)].as_ref())
        .split(area);

    let titles = RequestMenu::iter()
        .map(|item| Line::from(item.to_string()))
        .collect();

    let tabs = Tabs::new(titles)
        // .block(Block::default())
        .select(0)
        .style(Style::default().fg(Color::Cyan))
        .highlight_style(
            Style::default()
                .add_modifier(Modifier::BOLD)
                .bg(Color::Black),
        );

    frame.render_widget(tabs, chunks[0]);

    let block = Block::default()
        .title("request")
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded);

    let text = Paragraph::new(app.requestbar.body.clone())
        .block(block)
        .wrap(Wrap { trim: true })
        .alignment(Alignment::Left);

    frame.render_widget(text, chunks[1]);
}
