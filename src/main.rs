// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

extern crate anyhow;
#[macro_use]
extern crate clap;
extern crate cursive;

use anyhow::Result;
use cursive::align::HAlign;
use cursive::event::{Event, Key};
use cursive::theme::BaseColor::Black;
use cursive::theme::Color::Dark;
use cursive::theme::PaletteColor::{Background, Secondary};
use cursive::theme::{BorderStyle, Palette, Theme};
use cursive::traits::*;
use cursive::views::{Dialog, LinearLayout, PaddedView, SelectView, TextView};
use cursive::Cursive;

fn main() -> Result<()> {
    let m = clap_app!(myapp =>
        (version: "1.0")
        (author: "Magnus Aa. Hirth <magnus.hirth@gmail.com")
        (about: "Select between multiple arguments")
        (@arg VALS: ... +required "Values to select between")
    )
    .get_matches();

    let mut vals: Vec<String> = m.values_of("VALS").unwrap().map(String::from).collect();
    vals.sort();

    //
    // Application views: Filter label, filter text, and value list.
    //
    let label = TextView::new("Filter: ").style(Secondary);
    let filter = TextView::new("");
    let mut sel = SelectView::new().h_align(HAlign::Left);
    sel.add_all_str(&vals);
    sel.set_on_submit(|s, v: &String| {
        s.quit();
        print!("{}", v)
    });

    //
    // Global color palette for the application.
    //
    let mut palette = Palette::default();
    palette[Background] = Dark(Black);
    let theme = Theme {
        shadow: false,
        borders: BorderStyle::None,
        palette: palette,
    };

    let mut siv = cursive::default();

    //
    // Filter-callbacks for printable characters.
    //
    for ch in ('\x00'..'\x7f').filter(|c| !c.is_control()) {
        let vc = vals.clone();
        siv.set_global_callback(ch, move |s: &mut Cursive| {
            let flt = s
                .call_on_name("filter", |flt: &mut TextView| {
                    flt.append(ch);
                    flt.get_content().source().to_string()
                })
                .unwrap();
            s.call_on_name("list", |lst: &mut SelectView| {
                let vals: Vec<String> = vc.iter().filter(|s| valflt(s, &flt)).cloned().collect();
                lst.clear();
                lst.add_all_str(vals)
            });
        });
    }

    //
    // Filter-callbacks for removing text.
    // Backspace removes on character, delete removes all.
    //
    let vc = vals.clone();
    siv.set_global_callback(Event::Key(Key::Backspace), move |s: &mut Cursive| {
        let flt = s
            .call_on_name("filter", |flt: &mut TextView| {
                let mut content: String = flt.get_content().source().to_string();
                content.pop();
                flt.set_content(&content);
                content
            })
            .unwrap();
        s.call_on_name("list", |lst: &mut SelectView| {
            let vals: Vec<String> = vc.iter().filter(|s| valflt(s, &flt)).cloned().collect();
            lst.clear();
            lst.add_all_str(vals)
        });
    });
    let vc = vals.clone();
    siv.set_global_callback(Event::Key(Key::Del), move |s: &mut Cursive| {
        s.call_on_name("filter", |flt: &mut TextView| flt.set_content(""));
        s.call_on_name("list", |lst: &mut SelectView| {
            lst.clear();
            lst.add_all_str(&vc)
        });
    });

    //
    // Compose views and run the application.
    //
    siv.set_theme(theme);
    siv.add_layer(
        Dialog::around(
            LinearLayout::vertical()
                .child(PaddedView::lrtb(
                    0,
                    0,
                    0,
                    1,
                    LinearLayout::horizontal()
                        .child(label)
                        .child(filter.with_name("filter")),
                ))
                .child(sel.with_name("list")),
        )
        .min_width(20),
    );
    siv.run();
    Ok(())
}

///
/// Value filter, returning true if the pattern matches the given value.
///
fn valflt(value: &str, pattern: &str) -> bool {
    value.to_lowercase().contains(&pattern.to_lowercase())
}
