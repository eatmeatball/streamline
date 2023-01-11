// Generated from java-escape by ANTLR 4.11.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class FuncLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.11.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, FUNC=6, FUNCNAME=7, INT=8, FLOAT=9, 
		STRING=10, BOOL=11, ANY=12, MUL=13, DIV=14, ADD=15, SUB=16, NUMBER=17, 
		WHITESPACE=18, RETURN=19;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "FUNC", "FUNCNAME", "INT", "FLOAT", 
			"STRING", "BOOL", "ANY", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", 
			"RETURN"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'class'", "'func'", "'numberFunction'", 
			"'int'", "'float'", "'string'", "'bool'", "'any'", "'*'", "'/'", "'+'", 
			"'-'", null, null, "'return'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, "FUNC", "FUNCNAME", "INT", "FLOAT", 
			"STRING", "BOOL", "ANY", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", 
			"RETURN"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public FuncLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Func.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0013~\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002"+
		"\u0012\u0007\u0012\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u0010\u0004\u0010m\b\u0010\u000b\u0010"+
		"\f\u0010n\u0001\u0011\u0004\u0011r\b\u0011\u000b\u0011\f\u0011s\u0001"+
		"\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0000\u0000\u0013\u0001\u0001\u0003\u0002"+
		"\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013"+
		"\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011"+
		"#\u0012%\u0013\u0001\u0000\u0002\u0001\u000009\u0003\u0000\t\n\r\r  \u007f"+
		"\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000"+
		"\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000"+
		"\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000"+
		"\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011"+
		"\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015"+
		"\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019"+
		"\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d"+
		"\u0001\u0000\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001"+
		"\u0000\u0000\u0000\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000"+
		"\u0000\u0001\'\u0001\u0000\u0000\u0000\u0003)\u0001\u0000\u0000\u0000"+
		"\u0005+\u0001\u0000\u0000\u0000\u0007-\u0001\u0000\u0000\u0000\t/\u0001"+
		"\u0000\u0000\u0000\u000b5\u0001\u0000\u0000\u0000\r:\u0001\u0000\u0000"+
		"\u0000\u000fI\u0001\u0000\u0000\u0000\u0011M\u0001\u0000\u0000\u0000\u0013"+
		"S\u0001\u0000\u0000\u0000\u0015Z\u0001\u0000\u0000\u0000\u0017_\u0001"+
		"\u0000\u0000\u0000\u0019c\u0001\u0000\u0000\u0000\u001be\u0001\u0000\u0000"+
		"\u0000\u001dg\u0001\u0000\u0000\u0000\u001fi\u0001\u0000\u0000\u0000!"+
		"l\u0001\u0000\u0000\u0000#q\u0001\u0000\u0000\u0000%w\u0001\u0000\u0000"+
		"\u0000\'(\u0005(\u0000\u0000(\u0002\u0001\u0000\u0000\u0000)*\u0005)\u0000"+
		"\u0000*\u0004\u0001\u0000\u0000\u0000+,\u0005{\u0000\u0000,\u0006\u0001"+
		"\u0000\u0000\u0000-.\u0005}\u0000\u0000.\b\u0001\u0000\u0000\u0000/0\u0005"+
		"c\u0000\u000001\u0005l\u0000\u000012\u0005a\u0000\u000023\u0005s\u0000"+
		"\u000034\u0005s\u0000\u00004\n\u0001\u0000\u0000\u000056\u0005f\u0000"+
		"\u000067\u0005u\u0000\u000078\u0005n\u0000\u000089\u0005c\u0000\u0000"+
		"9\f\u0001\u0000\u0000\u0000:;\u0005n\u0000\u0000;<\u0005u\u0000\u0000"+
		"<=\u0005m\u0000\u0000=>\u0005b\u0000\u0000>?\u0005e\u0000\u0000?@\u0005"+
		"r\u0000\u0000@A\u0005F\u0000\u0000AB\u0005u\u0000\u0000BC\u0005n\u0000"+
		"\u0000CD\u0005c\u0000\u0000DE\u0005t\u0000\u0000EF\u0005i\u0000\u0000"+
		"FG\u0005o\u0000\u0000GH\u0005n\u0000\u0000H\u000e\u0001\u0000\u0000\u0000"+
		"IJ\u0005i\u0000\u0000JK\u0005n\u0000\u0000KL\u0005t\u0000\u0000L\u0010"+
		"\u0001\u0000\u0000\u0000MN\u0005f\u0000\u0000NO\u0005l\u0000\u0000OP\u0005"+
		"o\u0000\u0000PQ\u0005a\u0000\u0000QR\u0005t\u0000\u0000R\u0012\u0001\u0000"+
		"\u0000\u0000ST\u0005s\u0000\u0000TU\u0005t\u0000\u0000UV\u0005r\u0000"+
		"\u0000VW\u0005i\u0000\u0000WX\u0005n\u0000\u0000XY\u0005g\u0000\u0000"+
		"Y\u0014\u0001\u0000\u0000\u0000Z[\u0005b\u0000\u0000[\\\u0005o\u0000\u0000"+
		"\\]\u0005o\u0000\u0000]^\u0005l\u0000\u0000^\u0016\u0001\u0000\u0000\u0000"+
		"_`\u0005a\u0000\u0000`a\u0005n\u0000\u0000ab\u0005y\u0000\u0000b\u0018"+
		"\u0001\u0000\u0000\u0000cd\u0005*\u0000\u0000d\u001a\u0001\u0000\u0000"+
		"\u0000ef\u0005/\u0000\u0000f\u001c\u0001\u0000\u0000\u0000gh\u0005+\u0000"+
		"\u0000h\u001e\u0001\u0000\u0000\u0000ij\u0005-\u0000\u0000j \u0001\u0000"+
		"\u0000\u0000km\u0007\u0000\u0000\u0000lk\u0001\u0000\u0000\u0000mn\u0001"+
		"\u0000\u0000\u0000nl\u0001\u0000\u0000\u0000no\u0001\u0000\u0000\u0000"+
		"o\"\u0001\u0000\u0000\u0000pr\u0007\u0001\u0000\u0000qp\u0001\u0000\u0000"+
		"\u0000rs\u0001\u0000\u0000\u0000sq\u0001\u0000\u0000\u0000st\u0001\u0000"+
		"\u0000\u0000tu\u0001\u0000\u0000\u0000uv\u0006\u0011\u0000\u0000v$\u0001"+
		"\u0000\u0000\u0000wx\u0005r\u0000\u0000xy\u0005e\u0000\u0000yz\u0005t"+
		"\u0000\u0000z{\u0005u\u0000\u0000{|\u0005r\u0000\u0000|}\u0005n\u0000"+
		"\u0000}&\u0001\u0000\u0000\u0000\u0003\u0000ns\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}