# roll-the-dices

A simple personal project to catch up with Golang.
Currently only a bunch of scripts that make some queries to mySQL, some requests to an API and calculates some random results from dices.

## Features
Current accepted commands and parameters, returned from API requests:

### Parameters
**dice**: [`operator`] [`modifier`] `dice_number` d `die faces`
>Both _dice number_ and _die faces_ should be integers.  _Operator_ (*+*) should be indicated if there is more than one dice in the command. _Modifier_ already can be set to **h** to take the higher value of the roll. If the modifier is added with an integer, it will be the number of higher results from the roll
>
_Examples_:
- 1d20 (one die of twenty faces)
- h3d10 (the higher result of three dice of ten faces)
- 1d8+3d6 (the sum of one die of eight faces plus three dice of six faces)
- 3h6d10 (the three higher results of six dice of ten faces)

**bonus**: `operator` `bonus quantifier`
>_Operator_ (*+* or *-*) should be indicated always, even if the command didn't use _dice_ parameters
>
_Examples_: +7. +3, -10, +1+3-2

**tag**: `text`
>_Text_ could include any text but special symbols or _dice_ or _bonus_ nomenclature
>
_Examples_: Initiative, test, searching for someone, etc

**custom tag**: `:text`
>_Text_ could include any text but special symbols or _dice_ or _bonus_ nomenclature. This is only used in **agrupa** command meaning it is an enclosure for the previous _dice_ and _bonus_ parameters.

### Commands
**tira** [_dice_[_dice_[...]]] [_bonus_[_bonus_[...]]] [_tag_]
> Resolves _**dice**_ parameters, if any. Default value is "1d20". Then sum _**bonus**_ parameters to the result, if any. If _**tag**_ parameters are passed they will be included in the beginning of the roll, working as a title or tag.
> 
_Examples:_

| command | result |
| --- | --- |
| tira 1d10+7 | 1d10[7]+7= 14 |
| tira +2 | 1d20[12]+2 = 14 |
| tira 1d20+3 Initiative | Initiative: 1d20[2]+3= 5 |
| tira 2h6d10 Jump | Jump: 6d10[10,2,1,6,3,4]= 16 |

**agrupa** [_dice_[_dice_[...]]] [_bonus_[_bonus_[...]]] [_custom_ | _tag_]
> As _**tira**_ command but returning a riched MarkDown response with results distributed by _**dice**_ parameters or _**custom**_ tags
>
_Example:_

**command**: agrupa 1d10+2d6+7-1

**result**: `1d10[7]`: 7

`2d6[1,5]`: 6

_Bonus_ : 6

**Total: 19**


_Example with custom tags:_

**command**: agrupa 1d10+7:slashing+2d6:fire

**result**: _slashing(10)_ = `1d10[3]+7`

_fire(4)_ = `2d6[1,3]`

**v** [_bonus_[_bonus_[...]]] [_tag_]
> Resolves an advantage roll: Rolling two twenty-sided dice and taking the higher roll
> 
_Examples:_

| command | result |
| --- | --- |
| v +7 | 2d10[7,12]+7= 19 |
| v -1 | 2d20[12,19]-1 = 18 |
| v Initiative | Initiative: 1d20[2,7]= 7 |

**dv** [_bonus_[_bonus_[...]]] [_tag_]
> Resolves a disadvantage roll: Rolling two twenty-sided dice and taking the lower roll
> 
_Examples:_

| command | result |
| --- | --- |
| dv +7 | 2d10[7,12]+7= 14 |
| dv -1 | 2d20[12,19]-1 = 11 |
| dv Initiative | Initiative: 1d20[2,7]= 2 |

**help**
> Displays markdown text with help about the commands and how to use them. Currently only spanish available.

**t**
> Displays a button keyboard to allow user to enter basic commands without actually writing.

## Versioning
Current stable version: v0.5.0-beta
#### Changelog
| version |  notes |
| --- | --- |
| v0.5 | - Included _agrupa_ command and performance improvement |
| v0.4 | - Included _help_ and _t_ commands |
| v0.3 | - Included new roll modifier to allow rolling any die number and taking any best die number |

## License
This project is licensed under GPL-3.0, see [LICENSE](./LICENSE) file for details
